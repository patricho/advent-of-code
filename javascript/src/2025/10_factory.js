import { assert, prettyPrint, readFileLines } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 7, () => part1('../inputs/2025/10-test.txt'))
        assert('part 2 test', 33, () => part2('../inputs/2025/10-test.txt'))

        assert('part 1', 522, () => part1('../inputs/2025/10-input.txt'))
        assert('part 2', 18105, () => part2('../inputs/2025/10-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const lines = readFileLines(filename)

    let sum = 0

    lines.forEach((line) => {
        sum += processPart1Line(line)
    })

    return sum
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const lines = readFileLines(filename)

    let sum = 0

    lines.forEach((line) => {
        sum += processPart2Line(line)
    })

    return sum
}

/**
 * @param {string} line
 * @returns {number}
 */
function processPart1Line(line) {
    const targetString = line.substring(1, line.indexOf(']'))
    const buttonsString = line.substring(line.indexOf(']') + 2, line.indexOf('{') - 1)

    const from = 0
    const to = parseTargetState(targetString)

    const buttons = buttonsString.split(' ').map((b) => parseButtonState(targetString, b))

    const [steps, _] = bfsFindSmallestSum(from, to, buttons)

    return steps.length
}

/**
 * Translates "#..##.#" to `1001101`
 * @param {string} input
 * @returns {number}
 */
function parseTargetState(input) {
    return parseInt(input.replaceAll('.', '0').replaceAll('#', '1'), 2)
}

/**
 * Translates "(0,3,4)" to `10011`
 * @param {string} input
 * @param {string} button
 * @returns {number}
 */
function parseButtonState(input, button) {
    let state = '0'.repeat(input.length)

    button
        .substring(1, button.length - 1)
        .split(',')
        .forEach((s) => {
            const idx = parseInt(s)
            state = state.slice(0, idx) + '1' + state.slice(idx + 1)
        })

    return parseInt(state, 2)
}

/**
 * @param {number} from
 * @param {number} target
 * @param {number[]} buttons
 * @returns {[number[],number]}
 */
function bfsFindSmallestSum(from, target, buttons) {
    const queue = [{ currentState: from, history: [], usedMask: 0 }]

    const visitedMasks = new Set()
    visitedMasks.add(0)

    let checks = 0

    while (queue.length > 0) {
        checks++

        // Dequeue
        const { currentState, history, usedMask } = queue.shift()

        // And we're done
        if (currentState === target) {
            return [history, checks]
        }

        for (let i = 0; i < buttons.length; i++) {
            const button = buttons[i]

            // Check if this number (at index i) has already been used in this path. (1 << i)
            // creates a mask with only the i-th bit set. The bitwise AND checks if that bit is
            // already set in the usedMask
            if ((usedMask & (1 << i)) !== 0) continue

            // Calculate the new mask by setting the i-th bit (bitwise OR)
            const newMask = usedMask | (1 << i)

            if (visitedMasks.has(newMask)) continue

            visitedMasks.add(newMask)

            const newHistory = [...history, button]

            // Modify state using a XOR with the current button
            const nextState = currentState ^ button

            // Enqueue the new state
            queue.push({ currentState: nextState, history: newHistory, usedMask: newMask })
        }
    }

    // Empty return if queue empties and we never hit the target
    return [[], checks]
}

/**
 * @param {string} line
 */
function processPart2Line(line) {
    const [matrix, initialTargets, maxTarget, numRows, numCols] = parsePart2Line(line)

    // Working copies for the matrix
    let targets = [...initialTargets]

    const [pivotMapping, freeColumns] = gaussianElimination(matrix, targets, numRows, numCols)

    let minTotal = 999999999

    /**
     * Calculate the solution for the equations by using the provided values for free variables
     * @param {number[]} freeValues
     */
    function checkFree(freeValues) {
        let solution = new Array(numCols).fill(0)

        // Assign the free variables
        freeColumns.forEach((colIdx, i) => (solution[colIdx] = freeValues[i]))

        // Solve for pivot variables (back substitution)
        for (let i = pivotMapping.length - 1; i >= 0; i--) {
            let pCol = pivotMapping[i]

            // Calculate sum of known variables to the right
            let currentSum = 0
            for (let j = pCol + 1; j < numCols; j++) {
                currentSum += matrix[i][j] * solution[j]
            }

            let val = (targets[i] - currentSum) / matrix[i][pCol]

            // Allow for tiny floating point noise
            let rounded = Math.round(val)
            if (Math.abs(val - rounded) > 1e-5) return null
            if (rounded < 0) return null // Can't press a button a negative number of times

            solution[pCol] = rounded
        }

        return solution.reduce((a, b) => a + b, 0)
    }

    /**
     * Recursive testing of values for free variables to find the optimal ones
     * @param {number} idx
     * @param {number[]} currentFreeValues
     */
    function recursiveFreeSearch(idx, currentFreeValues) {
        // Base case: All free variables assigned
        if (idx === freeColumns.length) {
            const result = checkFree(currentFreeValues)
            if (result !== null) {
                minTotal = Math.min(minTotal, result)
            }
            return
        }

        // Recursive Step
        for (let free = 0; free <= maxTarget; free++) {
            currentFreeValues.push(free)
            recursiveFreeSearch(idx + 1, currentFreeValues)
            currentFreeValues.pop()
        }
    }

    recursiveFreeSearch(0, [])

    return minTotal
}

/**
 * @param {string} line
 * @returns {[number[][], number[], number, number, number]}
 */
function parsePart2Line(line) {
    const buttonMatches = [...line.matchAll(/\((.*?)\)/g)].map((m) => m[1])
    const targetMatch = line.match(/\{(.*?)\}/)[1]

    // Original targets (keep these for limit calculation)
    const initialTargets = targetMatch.split(',').map(Number)

    const numRows = initialTargets.length
    const numCols = buttonMatches.length

    let matrix = Array.from({ length: numRows }, () => Array(numCols).fill(0))
    buttonMatches.forEach((buttonStr, colIdx) => {
        buttonStr
            .split(',')
            .map(Number)
            .forEach((rowIdx) => {
                matrix[rowIdx][colIdx] = 1
            })
    })

    const maxTarget = Math.max(...initialTargets)

    return [matrix, initialTargets, maxTarget, numRows, numCols]
}

/**
 * @param {number[][]} matrix
 * @param {number[]} targets
 * @param {number} numRows
 * @param {number} numCols
 * @returns {[number[], number[]]}
 */
function gaussianElimination(matrix, targets, numRows, numCols) {
    let pivotMapping = []
    let row = 0
    for (let col = 0; col < numCols && row < numRows; col++) {
        let pivotRow = row

        // Find pivot
        while (pivotRow < numRows && Math.abs(matrix[pivotRow][col]) <= 0) pivotRow++

        // Free variable column
        if (pivotRow === numRows) {
            continue
        }

        // Swap
        ;[matrix[row], matrix[pivotRow]] = [matrix[pivotRow], matrix[row]]
        ;[targets[row], targets[pivotRow]] = [targets[pivotRow], targets[row]]

        // Eliminate
        for (let r = row + 1; r < numRows; r++) {
            if (Math.abs(matrix[r][col]) >= 0) {
                let factor = matrix[r][col] / matrix[row][col]
                targets[r] -= factor * targets[row]
                for (let c = col; c < numCols; c++) {
                    matrix[r][c] -= factor * matrix[row][c]
                }
            }
        }

        pivotMapping[row] = col
        row++
    }

    const pivotColsSet = new Set(pivotMapping)
    const freeColumns = Array.from({ length: numCols }, (_, i) => i).filter((i) => !pivotColsSet.has(i))

    return [pivotMapping, freeColumns]
}
