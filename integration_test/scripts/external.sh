#!/bin/bash

# Запускаем шелл в background с перенаправлением ввода-вывода через named pipes

# Создаем именованные каналы для взаимодействия с шеллом
PIPE_IN=$(mktemp -u)
PIPE_OUT=$(mktemp -u)
mkfifo "$PIPE_IN" "$PIPE_OUT"

# Файл для вывода результата echo
OUTPUT_FILE="results/output_app.txt"

# Запуск вашей программы (шелла), читающей команды из PIPE_IN и пишущей вывод в PIPE_OUT
go run ../cmd/main.go < "$PIPE_IN" > "$PIPE_OUT" &

# Отправляем команду echo в входной pipe
echo "grep a testfile.txt" > "$PIPE_IN"
grep a testfile.txt > results/output_bash.txt

# Закрываем входной канал - сигнализируем, что команд больше не будет
exec 3>"$PIPE_IN"
exec 3>&-

# Читаем из pipe с выводом команды и сохраняем в файл
cat "$PIPE_OUT" > "$OUTPUT_FILE"

# Удаляем именованные каналы
rm "$PIPE_IN" "$PIPE_OUT"

