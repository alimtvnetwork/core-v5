// aliases.go defines backward-compatible type aliases for all generic arg types.
//
// The generic types (e.g., FuncWrap[T], Three[T1, T2, T3]) are the primary types.
// These aliases instantiate them with 'any' for untyped usage, preserving the
// original API surface.
//
// Migration: Replace old type names with the *Any suffix variants.
//
//	args.FuncWrap   → args.FuncWrapAny   (or args.FuncWrap[any])
//	args.ThreeFunc  → args.ThreeFuncAny  (or args.ThreeFunc[any, any, any])
//	args.Holder     → args.HolderAny     (or args.Holder[any])
package args

// FuncWrapAny is the untyped version of FuncWrap[T].
// Use this as a drop-in replacement for the former non-generic FuncWrap.
type FuncWrapAny = FuncWrap[any]

// OneAny is the untyped version of One[T1].
type OneAny = One[any]

// OneFuncAny is the untyped version of OneFunc[T1].
type OneFuncAny = OneFunc[any]

// TwoAny is the untyped version of Two[T1, T2].
type TwoAny = Two[any, any]

// TwoFuncAny is the untyped version of TwoFunc[T1, T2].
type TwoFuncAny = TwoFunc[any, any]

// ThreeAny is the untyped version of Three[T1, T2, T3].
type ThreeAny = Three[any, any, any]

// ThreeFuncAny is the untyped version of ThreeFunc[T1, T2, T3].
type ThreeFuncAny = ThreeFunc[any, any, any]

// FourAny is the untyped version of Four[T1, T2, T3, T4].
type FourAny = Four[any, any, any, any]

// FourFuncAny is the untyped version of FourFunc[T1, T2, T3, T4].
type FourFuncAny = FourFunc[any, any, any, any]

// FiveAny is the untyped version of Five[T1, T2, T3, T4, T5].
type FiveAny = Five[any, any, any, any, any]

// FiveFuncAny is the untyped version of FiveFunc[T1, T2, T3, T4, T5].
type FiveFuncAny = FiveFunc[any, any, any, any, any]

// SixAny is the untyped version of Six[T1, T2, T3, T4, T5, T6].
type SixAny = Six[any, any, any, any, any, any]

// SixFuncAny is the untyped version of SixFunc[T1, T2, T3, T4, T5, T6].
type SixFuncAny = SixFunc[any, any, any, any, any, any]

// HolderAny is the untyped version of Holder[T].
type HolderAny = Holder[any]
