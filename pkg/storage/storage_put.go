package storage

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/pyroscope-io/pyroscope/pkg/storage/dimension"
	"github.com/pyroscope-io/pyroscope/pkg/storage/segment"
	"github.com/pyroscope-io/pyroscope/pkg/storage/tree"
)

func (s *Storage) Put(_ context.Context, pi *PutInput) error {
	// TODO: This is a pretty broad lock. We should find a way to make these locks more selective.
	s.putMutex.Lock()
	defer s.putMutex.Unlock()
	if s.hc.IsOutOfDiskSpace() {
		return errOutOfSpace
	}
	if pi.StartTime.Before(s.retentionPolicy().LowerTimeBoundary()) {
		return errRetention
	}

	s.metrics.putTotal.Inc()
	if id, ok := pi.Key.ProfileID(); ok {
		return s.exemplars.insert(pi.Key.AppName(), id, pi.Val, pi.EndTime)
	}

	s.logger.WithFields(logrus.Fields{
		"startTime":       pi.StartTime.String(),
		"endTime":         pi.EndTime.String(),
		"key":             pi.Key.Normalized(),
		"samples":         pi.Val.Samples(),
		"units":           pi.Units,
		"aggregationType": pi.AggregationType,
	}).Debug("storage.Put")

	sk := pi.Key.SegmentKey()

	for k, v := range pi.Key.Labels() {
		err := s.labels.Put(k, v)
		if err != nil {
			return fmt.Errorf("failed to put labels: %w", err)
		}
	}

	for k, v := range pi.Key.Labels() {
		key := k + ":" + v
		r, err := s.dimensions.GetOrCreate(key)
		if err != nil {
			s.logger.Errorf("dimensions cache for %v: %v", key, err)
			continue
		}
		r.(*dimension.Dimension).Insert([]byte(sk))
		s.dimensions.Put(key, r)
	}

	r, err := s.segments.GetOrCreate(sk)
	if err != nil {
		return fmt.Errorf("segments cache for %v: %w", sk, err)
	}

	st := r.(*segment.Segment)
	st.SetMetadata(pi.SpyName, pi.SampleRate, pi.Units, pi.AggregationType)
	samples := pi.Val.Samples()

	err = st.Put(pi.StartTime, pi.EndTime, samples, func(depth int, t time.Time, r *big.Rat, addons []segment.Addon) {
		tk := pi.Key.TreeKey(depth, t)
		res, err := s.trees.GetOrCreate(tk)
		if err != nil {
			s.logger.Errorf("trees cache for %v: %v", tk, err)
			return
		}
		cachedTree := res.(*tree.Tree)
		treeClone := pi.Val.Clone(r)
		for _, addon := range addons {
			if res, ok := s.trees.Lookup(pi.Key.TreeKey(addon.Depth, addon.T)); ok {
				ta := res.(*tree.Tree)
				ta.RLock()
				treeClone.Merge(ta)
				ta.RUnlock()
			}
		}
		cachedTree.Lock()
		cachedTree.Merge(treeClone)
		cachedTree.Unlock()
		s.trees.Put(tk, cachedTree)
	})
	if err != nil {
		return err
	}

	s.segments.Put(sk, st)
	return nil
}
