import { readFileLines, assert, prettyPrint } from '../util/index.js'

export default main

function main() {
    prettyPrint(() => {
        assert('part 1 test', 40, () => part1('../inputs/2025/08-test.txt', 10))
        // assert('part 2 test', ???, part2('../inputs/2025/08-test.txt'))

        assert('part 1', 244188, () => part1('../inputs/2025/08-input.txt', 1000))
        // assert('part 2', ???, part2('../inputs/2025/08-input.txt'))
    })
}

/**
 * @param {string} filename
 * @param {number} target
 * @returns {number}
 */
function part1(filename, target) {
    const lines = readFileLines(filename)
    const coords = {}

    console.log('a:', performance.now())

    lines.forEach((line) => {
        const cs = line.split(',').map((s) => parseInt(s))
        const coord = { x: cs[0], y: cs[1], z: cs[2] }
        const key = coord.x + '.' + coord.y + '.' + coord.z
        coords[key] = coord
    })

    let pairs = []

    console.log('b:', performance.now())

    for (const ka in coords) {
        for (const kb in coords) {
            if (ka === kb) continue
            const [dx, dy, dz] = [coords[ka].x - coords[kb].x, coords[ka].y - coords[kb].y, coords[ka].z - coords[kb].z]
            pairs.push({ first: ka, second: kb, euclidian: Math.sqrt(dx * dx + dy * dy + dz * dz) })
        }
    }

    console.log('c:', performance.now())

    pairs.sort((a, b) => a.euclidian - b.euclidian)

    console.log('d:', performance.now())

    let circuitno = 0
    const connected = new Set()

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
            const tomerge = [coords[pair.first].circuit, coords[pair.second].circuit, circuitno]
            for (const k in coords) {
                if (tomerge.some((n) => n == coords[k].circuit)) {
                    coords[k].circuit = circuitno
                }
            }
        } else {
            let nextcircuit = coords[pair.first].circuit ?? coords[pair.second].circuit ?? circuitno

            coords[pair.first].circuit = nextcircuit
            coords[pair.second].circuit = nextcircuit
        }

        connected.add(pair.first + pair.second)
        connected.add(pair.second + pair.first)

        if (circuitno >= target) break
    }

    const circuits = {}

    console.log('e:', performance.now())

    // Count/group by circuit
    for (const k in coords) {
        const c = coords[k]
        if (!c.circuit) continue
        circuits['c' + c.circuit] = parseInt(circuits['c' + c.circuit] ?? 0) + 1
    }

    console.log('f:', performance.now())

    // Sort by count, sum the 3 first
    const sum = Object.values(circuits)
        .sort((a, b) => b - a)
        .splice(0, 3)
        .reduce((acc, n) => acc * n, 1)

    console.log('g:', performance.now())

    return sum
}
