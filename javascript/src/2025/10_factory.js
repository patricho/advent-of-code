import { readFileString, assert, prettyPrint, isPointInPolygon, readFileLines } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 7, () => part1('../inputs/2025/10-test.txt'))
        // assert('part 2 test', 0, () => part2('../inputs/2025/10-test.txt'))

        assert('part 1', 522, () => part1('../inputs/2025/10-input.txt'))
        // assert('part 2', 0, () => part2('../inputs/2025/10-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const lines = readFileLines(filename)

    let sum = 0

    lines.forEach((line, i) => {
        sum += processLine(i, line)
    })

    return sum
}

/**
 * @param {number} no
 * @param {string} line
 * @returns {number}
 */
function processLine(no, line) {
    const targetString = line.substring(1, line.indexOf(']'))
    const buttonsString = line.substring(line.indexOf(']') + 2, line.indexOf('{') - 1)

    const from = 0
    const to = parseTargetState(targetString)

    const buttons = buttonsString.split(' ').map((b) => parseButtonState(targetString, b))

    const [steps, _] = bfsFindSmallestSum(from, to, buttons)

    // console.log(no, targetString, 'steps:', steps.length, steps)

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

            // Don't visit a combination of numbers that we've already found
            if (visitedMasks.has(newMask)) continue

            // Mark the new combination as visited
            visitedMasks.add(newMask)

            // Create the new history
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
