Package mute provides an encapsulation with no-op Print methods.

See [example](example_test.go),

```console
$ go test -run Example -v
=== RUN   TestExample
error immutable
error mutable
info done
--- PASS: TestExample (0.00s)
PASS
ok      github.com/platinasystems/mute  0.001s
$ go test -run Example -v --verbose
=== RUN   TestExample
error immutable
error mutable
info hello world
info done
--- PASS: TestExample (0.00s)
PASS
ok      github.com/platinasystems/mute  0.001s
$ go test -run Example -v --quiet
=== RUN   TestExample
error immutable
info done
--- PASS: TestExample (0.00s)
PASS
ok      github.com/platinasystems/mute  0.001s
```

---

*&copy; 2022 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE](LICENSE).*
