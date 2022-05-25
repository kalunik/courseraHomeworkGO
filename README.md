# Разработка веб-сервисов на Go

>Прохождение курса на Coursera от МФТИ, VK, ФРОО. 

## Список заданий
-[x] 1 неделя:
  - [задание](./hw1_dirTree/hw1.md): утилита `dirTree` 
  - [решение](./hw1_dirTree/main.go)


- [ ] 2 неделя:  
  - [задание](./hw2_signer/hw2.md): конвейер функций (`pipeline`)
  - [решение](./hw2_signer/signer.go)


- [ ] 3 неделя: 
  - [задание](./hw3_bench/hw3.md): оптимизация программы с использованием `pprof`


- [ ] 4 неделя:
  - [задание](./hw4_test_coverage/hw4.md): http-сервис и покрытие его тестами

<details>
<summary>Материалы для чтения</summary>

- Неделя 1. [_конспект лекции_](./info/lectureGo1.pdf)
  <details>
  <summary>Материалы для дополнительного чтения на английском:</summary>

  - https://golang.org/ref/spec - спецификация языка

  - https://golang.org/ref/mem - модель памяти го. на начальном этапе не надо, но знать полезно

  - https://golang.org/doc/code.html - про организацию кода. GOPATH и пакеты

  - https://golang.org/cmd/go/

  - https://blog.golang.org/strings

  - https://blog.golang.org/slices

  - https://blog.golang.org/go-slices-usage-and-internals

  - https://github.com/golang/go/wiki - вики го на гитхабе. очень много полезной информации

  - https://blog.golang.org/go-maps-in-action

  - https://blog.golang.org/organizing-go-code

  - https://golang.org/doc/effective_go.html - основной сборник тайного знания, сюда вы будуте обращатсья в первое время часто

  - https://github.com/golang/go/wiki/CodeReviewComments как ревьювить (и писать код). обязательно к прочтению

  - https://divan.github.io/posts/avoid_gotchas/ - материал аналогичный 50 оттенков гоhttps://research.swtch.com/interfaces

  - https://research.swtch.com/godata

  - http://jordanorelli.com/post/42369331748/function-types-in-go-golang

  - https://www.devdungeon.com/content/working-files-go - работа с файлами

  - http://www.golangprograms.com - много how-to касательно базовых вещей в go

  - http://yourbasic.org/golang/ - ещё большой набор how-to где можно получить углублённую информацию по всем базовым вещам. очень полезны http://yourbasic.org/golang/blueprint/

  - https://github.com/Workiva/go-datastructures

  - https://github.com/enocom/gopher-reading-list - большая подборка статей по многим темам ( не только данной лекции )

  - https://www.youtube.com/watch?v=MzTcsI6tn-0 - как организовать код

  - https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1 - статья на предыдущую тему
  </details>

  <details>
  <summary>Материалы для дополнительного чтения на русском:</summary>

  - https://habrahabr.ru/company/mailru/blog/314804/ - 50 оттенков го. обязательно к прочтению. многое оттуда мы ещё не проходили, но на будущее - имейте ввиду

  - https://habrahabr.ru/post/306914/ - Разбираемся в Go: пакет io

  - https://habrahabr.ru/post/272383/ - постулаты go. Маленькая статья об основными принципах языка

  - https://habrahabr.ru/company/mailru/blog/301036/ - лучшие практики go

  - https://habrahabr.ru/post/308198/ - организация кода в go

  - https://habrahabr.ru/post/339192/ - Зачем в Go амперсанд и звёздочка (& и *)

  - https://habrahabr.ru/post/325468/ - как не наступать на грабли в Go

  - https://habrahabr.ru/post/276981/ - Краш-курс по интерфейсам в Go

  - http://golang-book.ru
  </details>

  <details>
  <summary>Литература по го на русском языке:</summary>
  
  - Язык программирования Go, Алан А. А. Донован, Брайан У. Керниган
  
  - Go на практике, Matt Butcher, Мэтт Фарина Мэтт

  - Программирование на Go. Разработка приложений XXI века, Марк Саммерфильд
  </details>


- Неделя 2. [_конспект лекции_](./info/lectureGo2.pdf)
  <details>
  <summary>Материалы для дополнительного чтения на английском:</summary>

  - https://blog.golang.org/race-detector

  - https://blog.golang.org/pipelines

  - https://blog.golang.org/advanced-go-concurrency-patterns

  - https://blog.golang.org/go-concurrency-patterns-timing-out-and

  - https://talks.golang.org/2012/concurrency.slide#1

  - https://www.goinggo.net/2017/10/the-behavior-of-channels.html

  - http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/ - рассказ про оптимизацию воркер пула

  - http://www.tapirgames.com/blog/golang-channel

  - http://www.tapirgames.com/blog/golang-channel-closing

  - https://github.com/golang/go/wiki/CommonMistakes
  </details>

  <details>
  <summary>Видео:</summary>
  - https://www.youtube.com/watch?v=5buaPyJ0XeQ - классное выступление Dave Cheney про функции первого класса и использование их с горутинами, очень рекомендую, оно небольшое 

  - https://www.youtube.com/watch?v=f6kdp27TYZs - Google I/O 2012 - Go Concurrency Patterns - очень рекомендую

  - https://www.youtube.com/watch?v=rDRa23k70CU&list=PLDWZ5uzn69eyM81omhIZLzvRhTOXvpeX9&index=15 - ещё одно хорошее видео про паттерны конкуренции в го

  - https://www.youtube.com/watch?v=KAWeC9evbGM - видео Андрея Смирнова с конференции Highload - в нём вы можете получить более детальную информацию по теме вводного видео (методы обработки запросов и плюсы неблокирующего подхода), о том, что там творится на системном уровне. На русском, не про go
  </details>
  
  <details>
  <summary>На русском:</summary>

  - https://habrahabr.ru/post/141853/ - как работают горутины

  - https://habrahabr.ru/post/308070/ - как работают каналы

  - https://habrahabr.ru/post/333654/ - как работает планировщик ( https://rakyll.org/scheduler/ )

  - https://habrahabr.ru/post/271789/ - танцы с мютексами
  </details>

  <details>
  <summary>Книги:</summary>

  - Язык программирования Go, Алан А. А. Донован, Брайан У. Керниган

  - Go на практике, Matt Butcher, Мэтт Фарина Мэтт

  - Программирование на Go. Разработка приложений XXI века, Марк Саммерфильд
  </details>


- Неделя 3. [_конспект лекции_](./info/lectureGo3.pdf)
  <details>
  <summary>Рефлексия и кодогенерация:</summary>
  
  - https://blog.golang.org/laws-of-reflection

  - https://habrahabr.ru/post/269887/

  - https://golang.org/src/go/ast/example_test.go

  - https://github.com/golang/tools/blob/master/cmd/stringer/stringer.go

  - https://golang.org/pkg/reflect/

  - http://blog.burntsushi.net/type-parametric-functions-golang/

  - https://habrahabr.ru/post/269887/

  - https://medium.com/kokster/go-reflection-creating-objects-from-types-part-i-primitive-types-6119e3737f5d

  - https://medium.com/kokster/go-reflection-creating-objects-from-types-part-ii-composite-types-69a0e8134f20
  </details>

  <details>
  <summary>Производительность. Материалы на русском:</summary>

  - https://habrahabr.ru/company/badoo/blog/301990/

  - https://habrahabr.ru/company/badoo/blog/324682/

  - https://habrahabr.ru/company/badoo/blog/332636/

  - https://habrahabr.ru/company/mailru/blog/331784/ - статья про то как Почта@Mail.ru держит 3 миллиона вебсокет-соединений
  </details>

  <details>
  <summary>Производительность. Материалы на английском:</summary>

  - https://blog.golang.org/profiling-go-programs

  - https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/ - большая статья, посвященная go tool trace

  - https://www.goinggo.net/2017/05/language-mechanics-on-stacks-and-pointers.html

  - https://www.rzaluska.com/blog/important-go-interfaces/

  - https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/preview

  - https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/

  - https://lwn.net/Articles/250967/ - не про го, но знать полезно

  - https://github.com/golang/go/wiki/Performance - много про то что можно вытащить из pprof-а

  - https://golang.org/doc/gdb

  - https://about.sourcegraph.com/go/advanced-testing-in-go/

  - https://about.sourcegraph.com/go/generating-better-machine-code-with-ssa/

  - https://about.sourcegraph.com/go/evolutionary-optimization-peter-bourgon/

  - https://signalfx.com/blog/a-pattern-for-optimizing-go-2/

  - http://go-talks.appspot.com/github.com/davecheney/presentations/performance-without-the-event-loop.slide#1

  - https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

  - https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast - вообще в блоге Дейва очень много полезной инфы по го

  - https://github.com/dgryski/go-perfbook/blob/master/performance.md

  - https://www.youtube.com/watch?v=NS1hmEWv4Ac - Make your Go go faster! Optimising performance through reducing memory allocations + слайды https://fosdem.org/2018/schedule/event/faster/attachments/slides/2510/export/events/attachments/faster/slides/2510/BryanBorehamGoOptimisation.pdf

  - https://www.youtube.com/watch?v=N3PWzBeLX2M - Profiling and Optimizing Go

  - https://www.youtube.com/watch?v=Lxt8Vqn4JiQ - Golang UK Conference 2017 | Filippo Valsorda - Fighting latency: the CPU profiler is not your ally

  - https://www.youtube.com/watch?v=ydWFpcoYraU - Finding Memory Leaks in Go Programs

  - http://www.integralist.co.uk/posts/profiling-go/

  - https://bravenewgeek.com/so-you-wanna-go-fast/

  - https://medium.com/@val_deleplace/go-code-refactoring-the-23x-performance-hunt-156746b522f7
  </details>

  <details>
  <summary>Тесты:</summary>

  - https://blog.golang.org/cover - расширенная информация о go test -cover
  </details>

  <details>
  <summary>Полезные инструменты:</summary>

  - https://mholt.github.io/json-to-go - позволяет по json сформировать структуру на go, в которую он может быть распакован

  - https://github.com/mailru/easyjson - кодогенератор для json от mail.ru
  </details>


- Неделя 4. [_конспект лекции_](./info/lectureGo3.pdf)
  <details>
  <summary>Документация:</summary>

  - https://golang.org/pkg/net/http/
  </details>

  <details>
  <summary>Дополнительные материалы:</summary>

  - https://gowebexamples.github.io/ - примеры касательно разработки веба

  - https://golang.org/doc/articles/wiki/

  - https://astaxie.gitbooks.io/build-web-application-with-golang/

  - https://github.com/thewhitetulip/web-dev-golang-anti-textbook/

  - https://codegangsta.gitbooks.io/building-web-apps-with-go/content/

  - http://www.golangprograms.com/

  - http://marcio.io/2015/07/cheap-mapreduce-in-go/

  - https://www.rzaluska.com/blog/important-go-interfaces/

  - https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/ - про таймауты

  - http://polyglot.ninja/golang-making-http-requests/

  - http://tumregels.github.io/Network-Programming-with-Go/
  </details>

  <details>
  <summary>На русском:</summary>

  - https://habrahabr.ru/post/330512/ - Многопользовательская игра на Go через telnet - чисто сеть
  </details>

</details>
