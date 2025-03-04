# Items

Count: 18

## Item 1: Stmt 0, `ImportOfModule`

```js
import { upper } from "module";

```

- Hoisted
- Side effects

## Item 2: Stmt 0, `ImportBinding(0)`

```js
import { upper } from "module";

```

- Hoisted
- Declares: `upper`

## Item 3: Stmt 1, `VarDeclarator(0)`

```js
export let foobar = "foo";

```

- Declares: `foobar`
- Write: `foobar`

## Item 4: Stmt 2, `VarDeclarator(0)`

```js
export const foo = foobar;

```

- Declares: `foo`
- Reads: `foobar`
- Write: `foo`

## Item 5: Stmt 3, `VarDeclarator(0)`

```js
const bar = "bar";

```

- Declares: `bar`
- Write: `bar`

## Item 6: Stmt 4, `Normal`

```js
foobar += bar;

```

- Reads: `bar`
- Write: `foobar`

## Item 7: Stmt 5, `VarDeclarator(0)`

```js
let foobarCopy = foobar;

```

- Declares: `foobarCopy`
- Reads: `foobar`
- Write: `foobarCopy`

## Item 8: Stmt 6, `Normal`

```js
foobar += "foo";

```

- Write: `foobar`

## Item 9: Stmt 7, `Normal`

```js
console.log(foobarCopy);

```

- Side effects
- Reads: `console`, `foobarCopy`

## Item 10: Stmt 8, `Normal`

```js
foobarCopy += "Unused";

```

- Write: `foobarCopy`

## Item 11: Stmt 9, `Normal`

```js
function internal() {
    return upper(foobar);
}

```

- Hoisted
- Declares: `internal`
- Reads (eventual): `upper`, `foobar`

## Item 12: Stmt 10, `Normal`

```js
export function external1() {
    return internal() + foobar;
}

```

- Hoisted
- Declares: `external1`
- Reads (eventual): `internal`, `foobar`

## Item 13: Stmt 11, `Normal`

```js
export function external2() {
    foobar += ".";
}

```

- Hoisted
- Declares: `external2`
- Write (eventual): `foobar`

# Phase 1
```mermaid
graph TD
    Item1;
    Item2;
    Item3;
    Item4;
    Item5;
    Item6;
    Item7;
    Item8;
    Item9;
    Item10;
    Item11;
    Item12;
    Item13;
    Item14;
    Item14["ModuleEvaluation"];
    Item15;
    Item15["export foobar"];
    Item16;
    Item16["export foo"];
    Item17;
    Item17["export external1"];
    Item18;
    Item18["export external2"];
```
# Phase 2
```mermaid
graph TD
    Item1;
    Item2;
    Item3;
    Item4;
    Item5;
    Item6;
    Item7;
    Item8;
    Item9;
    Item10;
    Item11;
    Item12;
    Item13;
    Item14;
    Item14["ModuleEvaluation"];
    Item15;
    Item15["export foobar"];
    Item16;
    Item16["export foo"];
    Item17;
    Item17["export external1"];
    Item18;
    Item18["export external2"];
    Item4 --> Item3;
    Item6 --> Item5;
    Item6 -.-> Item4;
    Item7 --> Item3;
    Item7 --> Item6;
    Item8 -.-> Item4;
    Item8 -.-> Item7;
    Item9 --> Item7;
    Item9 --> Item1;
    Item9 -.-> Item2;
    Item9 -.-> Item3;
    Item9 -.-> Item6;
    Item9 -.-> Item8;
    Item9 -.-> Item4;
    Item9 -.-> Item11;
    Item10 -.-> Item9;
```
# Phase 3
```mermaid
graph TD
    Item1;
    Item2;
    Item3;
    Item4;
    Item5;
    Item6;
    Item7;
    Item8;
    Item9;
    Item10;
    Item11;
    Item12;
    Item13;
    Item14;
    Item14["ModuleEvaluation"];
    Item15;
    Item15["export foobar"];
    Item16;
    Item16["export foo"];
    Item17;
    Item17["export external1"];
    Item18;
    Item18["export external2"];
    Item4 --> Item3;
    Item6 --> Item5;
    Item6 -.-> Item4;
    Item7 --> Item3;
    Item7 --> Item6;
    Item8 -.-> Item4;
    Item8 -.-> Item7;
    Item9 --> Item7;
    Item9 --> Item1;
    Item9 -.-> Item2;
    Item9 -.-> Item3;
    Item9 -.-> Item6;
    Item9 -.-> Item8;
    Item9 -.-> Item4;
    Item9 -.-> Item11;
    Item10 -.-> Item9;
    Item11 --> Item2;
    Item11 --> Item3;
    Item11 --> Item6;
    Item11 --> Item8;
    Item12 --> Item11;
    Item12 --> Item3;
    Item12 --> Item6;
    Item12 --> Item8;
    Item13 -.-> Item4;
    Item13 -.-> Item7;
```
# Phase 4
```mermaid
graph TD
    Item1;
    Item2;
    Item3;
    Item4;
    Item5;
    Item6;
    Item7;
    Item8;
    Item9;
    Item10;
    Item11;
    Item12;
    Item13;
    Item14;
    Item14["ModuleEvaluation"];
    Item15;
    Item15["export foobar"];
    Item16;
    Item16["export foo"];
    Item17;
    Item17["export external1"];
    Item18;
    Item18["export external2"];
    Item4 --> Item3;
    Item6 --> Item5;
    Item6 -.-> Item4;
    Item7 --> Item3;
    Item7 --> Item6;
    Item8 -.-> Item4;
    Item8 -.-> Item7;
    Item9 --> Item7;
    Item9 --> Item1;
    Item9 -.-> Item2;
    Item9 -.-> Item3;
    Item9 -.-> Item6;
    Item9 -.-> Item8;
    Item9 -.-> Item4;
    Item9 -.-> Item11;
    Item10 -.-> Item9;
    Item11 --> Item2;
    Item11 --> Item3;
    Item11 --> Item6;
    Item11 --> Item8;
    Item12 --> Item11;
    Item12 --> Item3;
    Item12 --> Item6;
    Item12 --> Item8;
    Item13 -.-> Item4;
    Item13 -.-> Item7;
    Item14 --> Item1;
    Item14 --> Item9;
    Item15 --> Item3;
    Item15 --> Item6;
    Item15 --> Item8;
    Item16 --> Item4;
    Item17 --> Item12;
    Item18 --> Item13;
```
# Final
```mermaid
graph TD
    N0["Items: [ItemId(ModuleEvaluation), ItemId(0, ImportOfModule), ItemId(0, ImportBinding(0)), ItemId(7, Normal)]"];
    N1["Items: [ItemId(Export((Atom('foobar' type=inline), #0)))]"];
    N2["Items: [ItemId(Export((Atom('foo' type=inline), #0)))]"];
    N3["Items: [ItemId(Export((Atom('external1' type=dynamic), #0))), ItemId(10, Normal)]"];
    N4["Items: [ItemId(Export((Atom('external2' type=dynamic), #0))), ItemId(11, Normal)]"];
    N5["Items: [ItemId(1, VarDeclarator(0))]"];
    N6["Items: [ItemId(2, VarDeclarator(0))]"];
    N7["Items: [ItemId(3, VarDeclarator(0))]"];
    N8["Items: [ItemId(4, Normal)]"];
    N9["Items: [ItemId(5, VarDeclarator(0))]"];
    N10["Items: [ItemId(6, Normal)]"];
    N11["Items: [ItemId(0, ImportBinding(0)), ItemId(9, Normal)]"];
    N0 --> N9;
    N0 --> N11;
    N0 --> N5;
    N0 --> N8;
    N0 --> N10;
    N0 --> N6;
    N1 --> N5;
    N1 --> N8;
    N1 --> N10;
    N2 --> N6;
    N3 --> N11;
    N3 --> N5;
    N3 --> N8;
    N3 --> N10;
    N4 --> N6;
    N4 --> N9;
    N6 --> N5;
    N8 --> N7;
    N8 --> N6;
    N9 --> N5;
    N9 --> N8;
    N10 --> N6;
    N10 --> N9;
    N11 --> N5;
    N11 --> N8;
    N11 --> N10;
```
# Modules (dev)
## Part 0
```js
import { foobarCopy } from "entry.js" assert {
    __turbopack_chunk__: 9
};
import "entry.js" assert {
    __turbopack_chunk__: 11
};
import "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
import "entry.js" assert {
    __turbopack_chunk__: 10
};
import "entry.js" assert {
    __turbopack_chunk__: 6
};
"module evaluation";
import "module";
import { upper } from "module";
console.log(foobarCopy);

```
## Part 1
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
import "entry.js" assert {
    __turbopack_chunk__: 10
};
export { foobar };

```
## Part 2
```js
import { foo } from "entry.js" assert {
    __turbopack_chunk__: 6
};
export { foo };

```
## Part 3
```js
import { internal } from "entry.js" assert {
    __turbopack_chunk__: 11
};
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
import "entry.js" assert {
    __turbopack_chunk__: 10
};
export { external1 };
function external1() {
    return internal() + foobar;
}

```
## Part 4
```js
import "entry.js" assert {
    __turbopack_chunk__: 6
};
import "entry.js" assert {
    __turbopack_chunk__: 9
};
export { external2 };
function external2() {
    foobar += ".";
}

```
## Part 5
```js
let foobar = "foo";
export { foobar };

```
## Part 6
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
const foo = foobar;
export { foo };

```
## Part 7
```js
const bar = "bar";
export { bar };

```
## Part 8
```js
import { bar } from "entry.js" assert {
    __turbopack_chunk__: 7
};
import "entry.js" assert {
    __turbopack_chunk__: 6
};
foobar += bar;
export { foobar };

```
## Part 9
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
let foobarCopy = foobar;
export { foobarCopy };

```
## Part 10
```js
import "entry.js" assert {
    __turbopack_chunk__: 6
};
import "entry.js" assert {
    __turbopack_chunk__: 9
};
foobar += "foo";
export { foobar };

```
## Part 11
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
import "entry.js" assert {
    __turbopack_chunk__: 10
};
import { upper } from "module";
function internal() {
    return upper(foobar);
}

```
## Merged (module eval)
```js
import "module";
import { upper } from "module";
let foobar = "foo";
export { foobar };
const bar = "bar";
export { bar };
const foo = foobar;
export { foo };
foobar += bar;
export { foobar };
let foobarCopy = foobar;
export { foobarCopy };
import { upper } from "module";
foobar += "foo";
export { foobar };
function internal() {
    return upper(foobar);
}
"module evaluation";
console.log(foobarCopy);

```
# Modules (prod)
## Part 0
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 7
};
"module evaluation";
import "module";
let foobarCopy = foobar;
console.log(foobarCopy);
export { foobarCopy };

```
## Part 1
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 7
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
export { foobar };

```
## Part 2
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
export { foo };
const foo = foobar;
export { foo };

```
## Part 3
```js
import { foobar } from "entry.js" assert {
    __turbopack_chunk__: 5
};
import "entry.js" assert {
    __turbopack_chunk__: 7
};
import "entry.js" assert {
    __turbopack_chunk__: 8
};
export { external1 };
import { upper } from "module";
function internal() {
    return upper(foobar);
}
function external1() {
    return internal() + foobar;
}

```
## Part 4
```js
export { external2 };
function external2() {
    foobar += ".";
}

```
## Part 5
```js
let foobar = "foo";
export { foobar };

```
## Part 6
```js
const bar = "bar";
export { bar };

```
## Part 7
```js
import { bar } from "entry.js" assert {
    __turbopack_chunk__: 6
};
foobar += bar;
export { foobar };

```
## Part 8
```js
foobar += "foo";
export { foobar };

```
## Merged (module eval)
```js
import "module";
let foobar = "foo";
export { foobar };
const bar = "bar";
export { bar };
foobar += bar;
export { foobar };
"module evaluation";
let foobarCopy = foobar;
console.log(foobarCopy);
export { foobarCopy };

```
## Merged (external1)
```js
import { upper } from "module";
let foobar = "foo";
export { foobar };
const bar = "bar";
export { bar };
foobar += bar;
export { foobar };
foobar += "foo";
export { foobar };
export { external1 };
function internal() {
    return upper(foobar);
}
function external1() {
    return internal() + foobar;
}

```
## Merged (external1,external2)
```js
import { upper } from "module";
let foobar = "foo";
export { foobar };
const bar = "bar";
export { bar };
foobar += bar;
export { foobar };
foobar += "foo";
export { foobar };
export { external1 };
function internal() {
    return upper(foobar);
}
function external1() {
    return internal() + foobar;
}
export { external2 };
function external2() {
    foobar += ".";
}

```
