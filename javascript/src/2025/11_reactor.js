import { assert, prettyPrint, readFileLines } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 5, () => part1('../inputs/2025/11-test.txt'))
        assert('part 2 test', 2, () => part2('../inputs/2025/11-test2.txt'))

        assert('part 1', 472, () => part1('../inputs/2025/11-input.txt'))
        assert('part 2', 526811953334940, () => part2('../inputs/2025/11-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const OUT = 'out'

    const devices = {}
    readFileLines(filename).forEach((line) => {
        const arr = line.split(': ')
        devices[arr[0]] = arr[1].split(' ')
    })

    const recurse = (/** @type {string} */ key) => {
        if (key === OUT) {
            return 1
        }

        let sum = 0

        devices[key].forEach((/** @type {string} */ subkey) => {
            sum += recurse(subkey)
        })

        return sum
    }

    return recurse('you')
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const OUT = 'out'
    const DAC = 'dac'
    const FFT = 'fft'

    const devices = {}
    readFileLines(filename).forEach((line) => {
        const arr = line.split(': ')
        devices[arr[0]] = arr[1].split(' ')
    })

    const memo = {}

    /**
     * @param {string} key
     * @param {boolean} seendac
     * @param {boolean} seenfft
     * @returns {number}
     */
    const recurse = (key, seendac, seenfft) => {
        if (key === OUT) {
            return seendac && seenfft ? 1 : 0
        }

        const memokey = key + (seendac ? '1' : '0') + (seenfft ? '1' : '0')

        if (memo[memokey] != undefined) {
            return memo[memokey]
        }

        if (key === DAC) seendac = true
        if (key === FFT) seenfft = true

        let sum = 0

        devices[key].forEach((/** @type {string} */ subkey) => {
            sum += recurse(subkey, seendac, seenfft)
        })

        memo[memokey] = sum
        return sum
    }

    return recurse('svr', false, false)
}
