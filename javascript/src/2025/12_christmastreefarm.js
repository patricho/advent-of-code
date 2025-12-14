import { readFileString, assert, prettyPrint, isPointInPolygon, readFileLines } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 2, () => part1('../inputs/2025/12-test.txt'))
        // assert('part 2 test', 0, () => part2('../inputs/2025/12-test.txt'))

        assert('part 1', 497, () => part1('../inputs/2025/12-input.txt'))
        // assert('part 2', 0, () => part2('../inputs/2025/12-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    const lines = readFileLines(filename)

    const trees = lines
        .filter((line) => line.indexOf('x') >= 0)
        .map((line) => {
            const treeparts = line.split(': ')

            const dims = treeparts[0].split('x').map((s) => parseInt(s))

            const count = treeparts[1]
                .split(' ')
                .map((s) => parseInt(s))
                .reduce((acc, n) => acc + n, 0)

            return { w: dims[0], h: dims[1], count }
        })

    // Let's try assuming all packages are solid 3x3 squares
    return trees.filter((tree) => {
        const max = Math.floor(tree.w / 3) * Math.floor(tree.h / 3)
        return max >= tree.count
    }).length
}
