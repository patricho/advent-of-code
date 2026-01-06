import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 357, () => part1('../inputs/2025/03-test.txt'))
        assert('part 2 test', 3121910778619, () => part2('../inputs/2025/03-test.txt'))

        assert('part 1', 17229, () => part1('../inputs/2025/03-input.txt'))
        assert('part 2', 170520923035051, () => part2('../inputs/2025/03-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    return checkFile(filename, 2)
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    return checkFile(filename, 12)
}

/**
 * @param {string} filename
 * @param {number} targetLen
 * @returns {number}
 */
function checkFile(filename, targetLen) {
    const input = readFileLines(filename)

    let result = 0

    for (const line of input) {
        const n = checkRange(line, targetLen)
        result += n
    }

    return result
}

/**
 * @param {string} row
 * @param {number} targetLen
 * @returns {number}
 */
function checkRange(row, targetLen) {
    let startIdx = 0
    let result = ''

    let maxIdx = row.length - targetLen

    while (true) {
        let pick = '0'

        for (let i = startIdx; i <= maxIdx; i++) {
            const candidate = row[i]

            if (candidate > pick) {
                pick = candidate
                startIdx = i + 1
            }
        }

        result += pick

        if (result.length >= targetLen) {
            break
        }

        maxIdx++
    }

    return parseInt(result)
}
