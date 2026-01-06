import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

const reTwice = /^(\d+)\1$/
const reMultiple = /^(\d+)\1+$/

function main() {
    prettyPrint(() => {
        assert('part 1 test', 1227775554, () => part1('../inputs/2025/02-test.txt'))
        assert('part 2 test', 4174379265, () => part2('../inputs/2025/02-test.txt'))

        assert('part 1', 12586854255, () => part1('../inputs/2025/02-input.txt'))
        assert('part 2', 17298174201, () => part2('../inputs/2025/02-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    return check(filename, reTwice)
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    return check(filename, reMultiple)
}

/**
 * @param {string} filename
 * @param {RegExp} re
 * @returns {number}
 */
function check(filename, re) {
    let errsum = 0

    const input = readFileLines(filename)
    const ranges = input[0].split(',')

    for (const rng of ranges) {
        errsum += checkRange(rng, re)
    }

    return errsum
}

/**
 * @param {string} rng
 * @param {RegExp} re
 * @returns {number}
 */
function checkRange(rng, re) {
    let errsum = 0

    const arr = rng.split('-')

    const start = parseInt(arr[0])
    const end = parseInt(arr[1])

    for (let n = start; n <= end; n++) {
        const nstr = n.toString()
        if (re.test(nstr)) {
            errsum += n
        }
    }

    return errsum
}
