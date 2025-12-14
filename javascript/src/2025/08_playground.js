import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 40, () => part1('../inputs/2025/08-test.txt', 10))
        assert('part 2 test', 25272, () => part2('../inputs/2025/08-test.txt'))

        assert('part 1', 244188, () => part1('../inputs/2025/08-input.txt', 1000))
        assert('part 2', 8361881885, () => part2('../inputs/2025/08-input.txt'))
    })
}

/**
 * @param {string} filename
 * @param {number} target
 * @returns {number}
 */
function part1(filename, target) {
    return solve(filename, 1, target)
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    return solve(filename, 2, 0)
}

/**
 * @param {string} filename
 * @param {number} part
 * @param {number} target
 * @returns {number}
 */
function solve(filename, part, target) {
    const [coords, pairs] = parseInput(filename, true)
    const coordsNo = Object.keys(coords).length

    let circuitno = 0
    const connected = new Set()
    let circuits = {}
    let latest = { first: 0, second: 0 }

    const countCircuits = () => {
        circuits = {}
        for (const k in coords) {
            const c = coords[k]
            if (!c.circuit) continue
            circuits['c' + c.circuit] = parseInt(circuits['c' + c.circuit] ?? 0) + 1
        }
        return [Object.keys(circuits).length, Object.keys(circuits)[0]]
    }

    for (let i = 0; i < pairs.length; i++) {
        const pair = pairs[i]

        if (connected.has(pair.first + pair.second) || connected.has(pair.second + pair.first)) {
            // Already directly connected
            continue
        }

        circuitno++

        if (
            coords[pair.first].circuit &&
            coords[pair.second].circuit &&
            coords[pair.first].circuit !== coords[pair.second].circuit
        ) {
            // Related circuits merged into one
            const tomerge = [coords[pair.first].circuit, coords[pair.second].circuit]
            for (const k in coords) {
                if (tomerge.some((n) => n == coords[k].circuit)) {
                    coords[k].circuit = coords[pair.first].circuit
                }
            }
        } else {
            let nextcircuit = coords[pair.first].circuit || coords[pair.second].circuit || circuitno

            coords[pair.first].circuit = nextcircuit
            coords[pair.second].circuit = nextcircuit
        }

        connected.add(pair.first + pair.second)
        connected.add(pair.second + pair.first)

        if (part === 1 && circuitno >= target) break

        if (part === 2) {
            const [numberOfCircuits, firstCircuit] = countCircuits()

            if (numberOfCircuits === 1 && circuits[firstCircuit] === coordsNo) {
                latest = { first: coords[pair.first].x, second: coords[pair.second].x }
                break
            }
        }
    }

    if (part == 1) {
        countCircuits()

        // Sort by count, sum the 3 first
        const sum = Object.values(circuits)
            .sort((a, b) => b - a)
            .splice(0, 3)
            .reduce((acc, n) => acc * n, 1)

        return sum
    }

    return latest.first * latest.second
}

/**
 * @param {string} filename
 * @returns {[any,any[]]}
 */
function parseInput(filename, sort) {
    const lines = readFileLines(filename)
    const coords = {}

    lines.forEach((line) => {
        const cs = line.split(',').map((s) => parseInt(s))
        const coord = { x: cs[0], y: cs[1], z: cs[2] }
        const key = coord.x + '.' + coord.y + '.' + coord.z
        coords[key] = coord
    })

    let pairs = []

    for (const ka in coords) {
        for (const kb in coords) {
            if (ka === kb) continue
            const [dx, dy, dz] = [coords[ka].x - coords[kb].x, coords[ka].y - coords[kb].y, coords[ka].z - coords[kb].z]
            pairs.push({ first: ka, second: kb, distance: Math.sqrt(dx * dx + dy * dy + dz * dz) })
        }
    }

    if (sort) pairs.sort((a, b) => a.distance - b.distance)

    return [coords, pairs]
}
