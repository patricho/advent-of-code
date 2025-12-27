import { readFileNumberGrid, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 36, () => part1('../inputs/2024/10-test.txt'))
        assert('part 2 test', 81, () => part2('../inputs/2024/10-test.txt'))
        assert('part 2 test', 3, () => part2('../inputs/2024/10-test2.txt'))

        assert('part 1', 786, () => part1('../inputs/2024/10-input.txt'))
        assert('part 2', 1722, () => part2('../inputs/2024/10-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const [_, totalTops] = solve(filename)
    return totalTops
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const [totalPaths, _] = solve(filename)
    return totalPaths
}

/**
 * @param {string} filename
 * @returns {number[]}
 */
function solve(filename) {
    const grid = readFileNumberGrid(filename, '')
    const rows = grid.length
    const cols = grid[0].length

    // Find start positions
    const starts = []
    for (let y = 0; y < rows; y++) {
        for (let x = 0; x < cols; x++) {
            if (grid[y][x] === 0) {
                starts.push([y, x])
            }
        }
    }

    const dirs = [
        [1, 0],
        [0, 1],
        [-1, 0],
        [0, -1],
    ]

    let visited = new Set()

    /**
     * @param {any[]} pos
     * @param {number} posval
     */
    function takeStep(path, pos, posval) {
        let topsCount = 0
        let pathsCount = 0

        dirs.forEach((dir) => {
            const poskey = pos[0].toString() + pos[1].toString()
            const pathkey = path.map((p) => p[0] + ':' + p[1]).join(',')

            if (posval === 9) {
                if (!visited.has(poskey)) {
                    topsCount++
                    visited.add(poskey)
                }
                if (!visited.has(pathkey)) {
                    pathsCount++
                    visited.add(pathkey)
                }

                return
            }

            const newpos = [pos[0] + dir[0], pos[1] + dir[1]]
            if (newpos[0] < 0 || newpos[1] < 0 || newpos[0] >= rows || newpos[1] >= cols) {
                // Out of bounds
                return
            }

            const newposval = grid[newpos[0]][newpos[1]]

            if (newposval !== posval + 1) {
                // Invalid step
                return
            }

            const newpath = [...path, newpos]

            const [p, t] = takeStep(newpath, newpos, newposval)
            pathsCount += p
            topsCount += t
        })

        return [pathsCount, topsCount]
    }

    let totalPaths = 0
    let totalTops = 0

    starts.forEach((start) => {
        visited = new Set()

        const [pathsCount, topsCount] = takeStep([], start, grid[start[0]][start[1]])

        totalPaths += pathsCount
        totalTops += topsCount
    })

    return [totalPaths, totalTops]
}
