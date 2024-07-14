<template>
	<main id="calculator-page">
		<div class="mt-2">
			<h1>Calculator</h1>
			<p>This is the calculator page</p>
		</div>
		<div class="calculator">
      <input type="text" class="result"
               :value="result" readonly />
      <div class="buttons">
        <button class="number" 
                @click="handleClick('7')">7</button>
        <button class="number"
                @click="handleClick('8')">8</button>
        <button class="number"
                @click="handleClick('9')">9</button>
        <button class="operator"
                @click="handleOperatorClick('/')">/</button>

        <button class="number"
                @click="handleClick('4')">4</button>
        <button class="number" 
                @click="handleClick('5')">5</button>
        <button class="number"
                @click="handleClick('6')">6</button>
        <button class="operator"
                @click="handleOperatorClick('*')">*</button>

        <button class="number"
                @click="handleClick('1')">1</button>
        <button class="number"
                @click="handleClick('2')">2</button>
        <button class="number" 
                @click="handleClick('3')">3</button>
        <button class="operator"
                @click="handleOperatorClick('-')">-</button>

        <button class="number"
                @click="handleClick('0')">0</button>
        <button class="number"
                @click="handleClick('.')">.</button>
        <button class="number" 
                @click="handleClick('00')">00</button>

        <button class="operator" 
                @click="handleOperatorClick('+')">+</button>

        <button class="clear"
                @click="handleClear">C</button>
        <button class="clear"
                @click="handleClearEntry">CE</button>
        <button class="equal"
                @click="calculate()">=</button>
      </div>
    </div>
	</main>
</template>

<script setup>
import { ref } from 'vue'

const result = ref('')
const calculated = ref(false)

const handleClick = (value) => {
  if (calculated.value) {
    result.value = value
    calculated.value = false 
  } else {
    result.value += value
  }
}

const handleClear = () => {
  result.value = ''
  calculated.value = false 
}

const handleClearEntry = () => {
  if (result.value && result.value.length > 0) {
    result.value = result.value.slice(0, -1)
    if (result.value.length === 0) {
      handleClear()
    }
  } else {
    handleClear()
  }
}

const handleOperatorClick = (operator) => {
  if (/[+*/-]$/.test(result.value)) {
    result.value = result.value.slice(0, -1) + operator
  } else {
    result.value += operator
  }
  calculated.value = false 
}

const calculate = () => {
  try {
    let evaluatedResult = eval(result.value.replace(/(^|[^0-9])0+(\d+)/g, '$1$2'))
    if (evaluatedResult === Infinity || evaluatedResult === -Infinity) {
      throw new Error('Divide by zero error')
    }
    result.value = Number.isFinite(evaluatedResult) ? evaluatedResult : 'Error'
    calculated.value = true
  } catch (error) {
    if (error.message === 'Divide by zero error') {
      result.value = 'Error: Divide by zero'
    } else {
      result.value = 'Error'
    }
  }
}
</script>

<style>

	#calculator-page {
    font-family: Arial, sans-serif;
    text-align: center;
    margin-top: 40px;
  }
  
  .calculator {
    width: 300px;
    margin: 0 auto;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 20px;
  }
  
  .result {
    width: 90%;
    padding: 10px;
    font-size: 20px;
    text-align: right;
    margin-bottom: 10px;
  }
  
  .buttons {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    grid-gap: 10px;
  }
  
  button {
    padding: 15px;
    font-size: 18px;
    cursor: pointer;
    border: none;
    outline: none;
  }
  
  .number, .operator {
    background-color: #f0f0f0;
  }
  
  .clear, .equal {
    background-color: #22c55e;
    color: #fff;
  }
  
  #calculator-page button:hover {
    background-color: #ddd;
  }
  
  .equal {
    grid-column: span 2;
  }
</style>