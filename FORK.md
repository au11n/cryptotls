# cryptotls — форк crypto/tls с пулом буферов

Форк стандартного Go `crypto/tls` (Go 1.25.5) для устранения per-Read аллокаций.

## Изменения относительно upstream

### tls/conn.go

1. **Строка ~125**: Добавлено поле `atLeast atLeastReader` в struct `Conn`.
   Встраивает `atLeastReader` прямо в `Conn`, вместо аллокации на хипе при каждом вызове `readFromUntil`.

2. **Функция `readFromUntil`**: Вместо `&atLeastReader{r, int64(needs)}` (heap escape через io.Reader interface) используется `&c.atLeast` — указатель на встроенное поле.

Эти два изменения устраняют ~2 аллокации на каждый TLS record (~44K allocs при ~22K records).

### Поиск изменений

Все изменения помечены комментарием `// FORK:`.

```bash
grep -rn 'FORK:' tls/
```

## Стабы для internal-пакетов

Стандартная библиотека использует `internal/` пакеты, недоступные из пользовательского кода. Замены:

| Оригинал | Замена | Описание |
|----------|--------|----------|
| `internal/godebug` | `stubs/godebug` | No-op, всегда возвращает "" |
| `internal/cpu` | `stubs/cpu` | Hardcoded CPU features для amd64 |
| `internal/goarch` | `stubs/goarch` | Константы для amd64 |
| `internal/byteorder` | `stubs/byteorder` | Обёртка над encoding/binary |
| `internal/stringslite` | `stubs/stringslite` | Обёртка над strings |
| `crypto/tls/internal/fips140tls` | `tls/internal/fips140tls` | No-op, FIPS не требуется |
| `crypto/internal/fips140/check` | `internal/fips140/check` | No-op |
| `crypto/internal/fips140` (indicator, cast) | `internal/fips140` | Упрощённые стабы |
| `crypto/internal/fips140hash` | `internal/fips140hash` | Passthrough без unwrap |
| `crypto/internal/sysrand` | `internal/sysrand` | Прямой syscall вместо internal/syscall/unix |

## При обновлении Go

1. Скопировать файлы из нового `crypto/tls/` в `tls/`
2. Скопировать `crypto/internal/` в `internal/`
3. Заменить imports: `sed -i 's|"crypto/internal/|"cryptotls/internal/|g'` и т.д.
4. Заново применить изменения из секции "Изменения" (искать по `// FORK:`)
5. Проверить, не появились ли новые `internal/` зависимости
