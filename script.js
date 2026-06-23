document.addEventListener('DOMContentLoaded', () => {
    const displayElement = document.getElementById('display');
    const memoryIndicator = document.querySelector('.memory-indicator');

    let currentInput = '0';
    let previousInput = null;
    let operator = null;
    let memory = 0;
    let resetDisplay = false;

    const updateDisplay = () => {
        // Limit display length to avoid overflow, though CSS handles scrolling
        displayElement.textContent = currentInput.length > 12 ? currentInput.slice(0, 12) : currentInput;
    };

    const handleNumber = (num) => {
        if (currentInput === '0' || resetDisplay) {
            currentInput = num;
            resetDisplay = false;
        } else {
            currentInput += num;
        }
        updateDisplay();
    };

    const handleDecimal = () => {
        if (resetDisplay) {
            currentInput = '0.';
            resetDisplay = false;
        } else if (!currentInput.includes('.')) {
            currentInput += '.';
        }
        updateDisplay();
    };

    const handleOperator = (op) => {
        if (operator !== null && !resetDisplay) {
            calculate();
        }
        previousInput = currentInput;
        operator = op;
        resetDisplay = true;
    };

    const calculate = () => {
        if (operator === null || previousInput === null) return;

        let result;
        const prev = parseFloat(previousInput);
        const current = parseFloat(currentInput);

        switch (operator) {
            case '+':
                result = prev + current;
                break;
            case '-':
                result = prev - current;
                break;
            case '*':
                result = prev * current;
                break;
            case '/':
                result = prev / current;
                break;
            default:
                return;
        }

        currentInput = String(result);
        operator = null;
        previousInput = null;
        updateDisplay();
        resetDisplay = true;
    };

    const handlePercentage = () => {
        currentInput = String(parseFloat(currentInput) / 100);
        updateDisplay();
    };

    const handleSqrt = () => {
        currentInput = String(Math.sqrt(parseFloat(currentInput)));
        updateDisplay();
    };

    const handleClearEntry = () => {
        currentInput = '0';
        updateDisplay();
    };

    const handleAllClear = () => {
        currentInput = '0';
        previousInput = null;
        operator = null;
        updateDisplay();
    };

    const handleBackspace = () => {
        if (currentInput.length > 1) {
            currentInput = currentInput.slice(0, -1);
        } else {
            currentInput = '0';
        }
        updateDisplay();
    };

    const handleMemoryRecall = () => {
        currentInput = String(memory);
        updateDisplay();
        resetDisplay = true;
    };

    const handleMemoryPlus = () => {
        memory += parseFloat(currentInput);
        showMemoryIndicator();
    };

    const handleMemoryMinus = () => {
        memory -= parseFloat(currentInput);
        showMemoryIndicator();
    };

    const showMemoryIndicator = () => {
        // Just show MEMORY text if memory is not 0
        if (memory !== 0) {
            memoryIndicator.style.visibility = 'visible';
            memoryIndicator.classList.add('visible');
        } else {
            memoryIndicator.style.visibility = 'hidden';
            memoryIndicator.classList.remove('visible');
        }
    };

    document.querySelectorAll('.btn').forEach(button => {
        button.addEventListener('click', () => {
            const action = button.dataset.action;
            const num = button.dataset.num;
            const op = button.dataset.op;

            if (num !== undefined) {
                handleNumber(num);
            } else if (op !== undefined) {
                handleOperator(op);
            } else if (action === 'decimal') {
                handleDecimal();
            } else if (action === 'calculate') {
                calculate();
            } else if (action === 'all-clear') {
                handleAllClear();
            } else if (action === 'clear-entry') {
                handleClearEntry();
            } else if (action === 'backspace') {
                handleBackspace();
            } else if (action === 'percent') {
                handlePercentage();
            } else if (action === 'sqrt') {
                handleSqrt();
            } else if (action === 'memory-recall') {
                handleMemoryRecall();
            } else if (action === 'memory-plus') {
                handleMemoryPlus();
            } else if (action === 'memory-minus') {
                handleMemoryMinus();
            }
        });
    });
});
