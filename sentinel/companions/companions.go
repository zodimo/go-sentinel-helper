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

type structCompanion[T any] struct {
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

func NewStructCompanion[T any](
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
) StructCompanion[T] {
	return &structCompanion[T]{
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

func (sc *structCompanion[T]) Unspecified() T {
	return sc.unspecified()
}
func (sc *structCompanion[T]) String(s T) string {
	return sc.stringer(s)
}
func (sc *structCompanion[T]) IsSpecified(s T) bool {
	return sc.isSpecified(s)
}
func (sc *structCompanion[T]) IsUnspecified(s T) bool {
	return sc.isUnspecified(s)
}
func (sc *structCompanion[T]) TakeOrElse(src, def T) T {
	return sc.takeOrElse(src, def)
}
func (sc *structCompanion[T]) Merge(a, b T) T {
	return sc.merge(a, b)
}
func (sc *structCompanion[T]) Same(a, b T) bool {
	return sc.same(a, b)
}
func (sc *structCompanion[T]) SemanticEqual(a, b T) bool {
	return sc.semanticEqual(a, b)
}
func (sc *structCompanion[T]) Equal(a, b T) bool {
	return sc.equal(a, b)
}

func (sc *structCompanion[T]) Coalesce(ptr, def T) T {
	return sc.coalesce(ptr, def)
}
func (sc *structCompanion[T]) Copy(s T, options ...func(T)) T {
	return sc.copy(s, options...)
}
