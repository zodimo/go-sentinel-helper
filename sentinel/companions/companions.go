package companions

type scalarCompanion[T any] struct {
	unspecified   func() T
	stringer      func(T) string
	isSpecified   func(T) bool
	isUnspecified func(T) bool
	takeOrElse    func(src, def T) T
	merge         func(a, b T) T
	same          func(a, b T) bool
	semanticEqual func(a, b T) bool
	equal         func(a, b T) bool
	copy          func(s T) T
}

func NewScalarCompanion[T any](
	unspecified func() T,
	stringer func(T) string,
	isSpecified func(T) bool,
	isUnspecified func(T) bool,
	takeOrElse func(src, def T) T,
	merge func(a, b T) T,
	same func(a, b T) bool,
	semanticEqual func(a, b T) bool,
	equal func(a, b T) bool,
	copy func(s T) T,
) ScalarCompanion[T] {
	return &scalarCompanion[T]{
		unspecified:   unspecified,
		stringer:      stringer,
		isSpecified:   isSpecified,
		isUnspecified: isUnspecified,
		takeOrElse:    takeOrElse,
		merge:         merge,
		same:          same,
		semanticEqual: semanticEqual,
		equal:         equal,
		copy:          copy,
	}
}

func (sc *scalarCompanion[T]) Unspecified() T {
	return sc.unspecified()
}
func (sc *scalarCompanion[T]) String(s T) string {
	return sc.stringer(s)
}
func (sc *scalarCompanion[T]) IsSpecified(s T) bool {
	return sc.isSpecified(s)
}
func (sc *scalarCompanion[T]) IsUnspecified(s T) bool {
	return sc.isUnspecified(s)
}
func (sc *scalarCompanion[T]) TakeOrElse(src, def T) T {
	return sc.takeOrElse(src, def)
}
func (sc *scalarCompanion[T]) Merge(a, b T) T {
	return sc.merge(a, b)
}
func (sc *scalarCompanion[T]) Same(a, b T) bool {
	return sc.same(a, b)
}
func (sc *scalarCompanion[T]) SemanticEqual(a, b T) bool {
	return sc.semanticEqual(a, b)
}
func (sc *scalarCompanion[T]) Equal(a, b T) bool {
	return sc.equal(a, b)
}
func (sc *scalarCompanion[T]) Copy(s T) T {
	return sc.copy(s)
}

type structCompoanion[T any] struct {
	unspecified   func() T
	stringer      func(T) string
	isSpecified   func(T) bool
	isUnspecified func(T) bool
	takeOrElse    func(src, def T) T
	merge         func(a, b T) T
	same          func(a, b T) bool
	semanticEqual func(a, b T) bool
	equal         func(a, b T) bool
	coalesce      func(ptr, def T) T
	copy          func(s T, options ...func(T)) T
}

func NewStructCompoanion[T any](
	unspecified func() T,
	stringer func(T) string,
	isSpecified func(T) bool,
	isUnspecified func(T) bool,
	takeOrElse func(src, def T) T,
	merge func(a, b T) T,
	same func(a, b T) bool,
	semanticEqual func(a, b T) bool,
	equal func(a, b T) bool,
	coalesce func(ptr, def T) T,
	copy func(s T, options ...func(T)) T,
) StructCompoanion[T] {
	return &structCompoanion[T]{
		unspecified:   unspecified,
		stringer:      stringer,
		isSpecified:   isSpecified,
		isUnspecified: isUnspecified,
		takeOrElse:    takeOrElse,
		merge:         merge,
		same:          same,
		semanticEqual: semanticEqual,
		equal:         equal,
		coalesce:      coalesce,
		copy:          copy,
	}
}

func (sc *structCompoanion[T]) Unspecified() T {
	return sc.unspecified()
}
func (sc *structCompoanion[T]) String(s T) string {
	return sc.stringer(s)
}
func (sc *structCompoanion[T]) IsSpecified(s T) bool {
	return sc.isSpecified(s)
}
func (sc *structCompoanion[T]) IsUnspecified(s T) bool {
	return sc.isUnspecified(s)
}
func (sc *structCompoanion[T]) TakeOrElse(src, def T) T {
	return sc.takeOrElse(src, def)
}
func (sc *structCompoanion[T]) Merge(a, b T) T {
	return sc.merge(a, b)
}
func (sc *structCompoanion[T]) Same(a, b T) bool {
	return sc.same(a, b)
}
func (sc *structCompoanion[T]) SemanticEqual(a, b T) bool {
	return sc.semanticEqual(a, b)
}
func (sc *structCompoanion[T]) Equal(a, b T) bool {
	return sc.equal(a, b)
}

func (sc *structCompoanion[T]) Coalesce(ptr, def T) T {
	return sc.coalesce(ptr, def)
}
func (sc *structCompoanion[T]) Copy(s T, options ...func(T)) T {
	return sc.Copy(s, options...)
}
