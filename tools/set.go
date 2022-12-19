package tools

type Set[T comparable] map[T]bool

func GetSet[T comparable]() Set[T] {
	return make(Set[T])
}

func GetSetFrom[T comparable](keys []T) Set[T] {
	s := GetSet[T]()
	for _, k := range keys {
		s.Add(k)
	}
	return s
}

func (s Set[T]) Add(key T) {
	s[key] = true
}

func (s Set[T]) Delete(key T) {
	delete(s, key)
}

func (s Set[T]) SoftDelete(key T) {
	s[key] = false
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Contains(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) GetKeys() []T {
	res := []T{}

	for k := range s {
		res = append(res, k)
	}

	return res
}

func (s Set[T]) GetIntersetion(otherSet Set[T]) Set[T] {
	intersection := GetSet[T]()

	if s.IsEmpty() || otherSet.IsEmpty() {
		return intersection
	}

	biggerSet := &s
	tinierSet := &otherSet

	if len(s) < len(otherSet) {
		biggerSet = &otherSet
		tinierSet = &s
	}

	for key := range *biggerSet {
		if tinierSet.Contains(key) {
			intersection.Add(key)
		}
	}

	return intersection
}

func (s Set[T]) Clone() Set[T] {
	clone := GetSet[T]()

	for key := range s {
		clone.Add(key)
	}

	return clone
}

func (s Set[T]) GetUnion(otherSet Set[T]) Set[T] {
	unionRes := GetSet[T]()

	for key := range s {
		unionRes.Add(key)
	}
	for key := range otherSet {
		unionRes.Add(key)
	}

	return unionRes
}

func (s Set[T]) Equals(otherSet Set[T]) bool {
	if len(s) != len(otherSet) {
		return false
	}
	for key := range s {
		if !otherSet.Contains(key) {
			return false
		}
	}

	return true
}
