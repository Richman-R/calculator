package main

import (
    "calculator/internal/usecase"
    "calculator/internal/delivery/console"
)

func main() {
    // 1. Создаем калькулятор (usecase)
    calc := usecase.NewCalculator()
    
    // 2. Создаем обработчик (delivery), передаем ему калькулятор
    handler := console.NewHandler(calc)
    
    // 3. Запускаем REPL-цикл
    handler.Run()
}