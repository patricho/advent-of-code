import { readFileLines, assert, prettyPrint, DIRECTIONS_AND_DIAGONALS, outOfBounds } from '../util/index.js'

export default main

//const ON = '@'
const OFF = '.'
const MARK = 'X'

/** @type {string[][]} */
let grid = []

function main() {
    prettyPrint(() => {
        assert('part 1 test', 13, () => part1('../inputs/2025/04-test.txt'))
        assert('part 2 test', 43, () => part2('../inputs/2025/04-test.txt'))

        assert('part 1', 1491, () => part1('../inputs/2025/04-input.txt'))
        assert('part 2', 8722, () => part2('../inputs/2025/04-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    initGrid(filename)
    return countNeighbors()
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    let result = 0

    initGrid(filename)

    while (true) {
        const round = countNeighbors()
        result += round

        if (round === 0) {
            break
        }

        removeNeighbors()
    }

    return result
}

/**
 * @param {string} filename
 */
function initGrid(filename) {
    const lines = readFileLines(filename)
    grid = lines.map((line) => line.split(''))
}

/**
 * @returns {number}
 */
function countNeighbors() {
    let result = 0

    for (let y = 0; y < grid.length; y++) {
        for (let x = 0; x < grid[y].length; x++) {
            if (grid[y][x] === OFF) {
                continue
            }

            let neighbors = 0

            for (const dir of DIRECTIONS_AND_DIAGONALS) {
                const nx = x + dir.x
                const ny = y + dir.y

                if (outOfBounds(grid, ny, nx)) {
                    continue
                }

                if (grid[ny][nx] !== OFF) {
                    neighbors++
                }
            }

            if (neighbors < 4) {
                grid[y][x] = MARK
                result++
            }
        }
    }

    return result
}

function removeNeighbors() {
    for (let y = 0; y < grid.length; y++) {
        for (let x = 0; x < grid[y].length; x++) {
            if (grid[y][x] === MARK) {
                grid[y][x] = OFF
            }
        }
    }
}
