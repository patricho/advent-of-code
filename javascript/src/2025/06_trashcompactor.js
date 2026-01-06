import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 4277556, () => part1('../inputs/2025/06-test.txt'))
        assert('part 2 test', 3263827, () => part2('../inputs/2025/06-test.txt'))

        assert('part 1', 5060053676136, () => part1('../inputs/2025/06-input.txt'))
        assert('part 2', 9695042567249, () => part2('../inputs/2025/06-input.txt'))
    })
}

/**
 * @typedef {Object} Equation
 * @property {number[]} numbers
 * @property {string} sign
 */

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const eqs = parseInput(filename)
    return sumEquations(eqs)
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const lines = readFileLines(filename)
    const columns = lines[0].length

    /** @type {Equation[]} */
    const equations = []

    /** @type {Equation} */
    let pending = { numbers: [], sign: '' }

    for (let x = columns - 1; x >= 0; x--) {
        let empty = true
        let numstring = ''

        for (const line of lines) {
            const n = line[x]
            if (n === ' ') {
                continue
            }

            if (n === '*' || n === '+') {
                pending.sign = n
            } else {
                numstring += n
            }
            empty = false
        }

        if (empty) {
            equations.push(pending)
            pending = { numbers: [], sign: '' }
        } else {
            pending.numbers.push(parseInt(numstring))
        }
    }

    if (pending.numbers.length > 0) {
        equations.push(pending)
    }

    return sumEquations(equations)
}

/**
 * @param {string} filename
 * @returns {Equation[]}
 */
function parseInput(filename) {
    const lines = readFileLines(filename)

    const lastline = lines[lines.length - 1]

    const columns = lastline.trim().split(/\s+/)

    /** @type {Equation[]} */
    const equations = columns.map(() => ({ numbers: [], sign: '' }))

    for (let i = 0; i < lines.length - 1; i++) {
        const parts = lines[i].trim().split(/\s+/)
        for (let j = 0; j < parts.length; j++) {
            equations[j].numbers.push(parseInt(parts[j]))
        }
    }

    const lastLineParts = lastline.trim().split(/\s+/)

    for (let j = 0; j < columns.length; j++) {
        equations[j].sign = lastLineParts[j][0]
    }

    return equations
}

/**
 * @param {Equation[]} equations
 * @returns {number}
 */
function sumEquations(equations) {
    let result = 0

    for (const eq of equations) {
        let res = eq.numbers[0]

        switch (eq.sign) {
            case '*':
                for (let i = 1; i < eq.numbers.length; i++) {
                    res *= eq.numbers[i]
                }
                break
            case '+':
                for (let i = 1; i < eq.numbers.length; i++) {
                    res += eq.numbers[i]
                }
                break
        }

        result += res
    }

    return result
}
