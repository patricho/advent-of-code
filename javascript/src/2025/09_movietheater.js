import { readFileLines, assert, prettyPrint, isPointInPolygon } from '../util/index.js'

export default main

let coords = []
let pointsInPolygon = new Set()
let pointsOutsidePolygon = new Set()

function main() {
    prettyPrint(() => {
        assert('part 1 test', 50, () => part1('../inputs/2025/09-test.txt'))
        assert('part 2 test', 24, () => part2('../inputs/2025/09-test.txt'))

        assert('part 1', 4735268538, () => part1('../inputs/2025/09-input.txt'))
        assert('part 2', 1537458069, () => part2('../inputs/2025/09-input.txt'))
    })
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part1(filename) {
    coords = readFileLines(filename).map((line) => line.split(',').map((s) => parseInt(s)))

    let maxArea = -1

    coords.forEach((a) => {
        coords.forEach((b) => {
            const area = (Math.abs(b[0] - a[0]) + 1) * (Math.abs(b[1] - a[1]) + 1)
            if (area > maxArea) maxArea = area
        })
    })

    return maxArea
}

/**
 * @param {string} filename
 * @returns {number}
 */
function part2(filename) {
    const coordsSet = new Set()

    coords = readFileLines(filename).map((line) => {
        coordsSet.add(line)
        const nums = line.split(',').map((s) => parseInt(s))
        return { x: nums[0], y: nums[1] }
    })

    let maxArea = 0

    coords.forEach((a) => {
        coords.forEach((b) => {
            if (a.x === b.x && a.y === b.y) return

            const minX = Math.min(a.x, b.x)
            const maxX = Math.max(a.x, b.x)
            const minY = Math.min(a.y, b.y)
            const maxY = Math.max(a.y, b.y)

            const area = (Math.abs(maxX - minX) + 1) * (Math.abs(maxY - minY) + 1)
            if (area < maxArea) return

            let enclosed = true

            // Early exit if any of the corners are outside the polygon
            if (
                !inPolygon({ x: minX, y: minY }, coords) ||
                !inPolygon({ x: minX, y: maxY }, coords) ||
                !inPolygon({ x: maxX, y: minY }, coords) ||
                !inPolygon({ x: maxX, y: maxY }, coords)
            ) {
                return
            }

            // Check along both Y edges
            for (let y = minY; y <= maxY; y++) {
                if (!inPolygon({ x: minX, y }, coords) || !inPolygon({ x: maxX, y }, coords)) {
                    enclosed = false
                    break
                }
            }

            if (!enclosed) return

            // Check along both X edges
            for (let x = minX; x <= maxX; x++) {
                if (!inPolygon({ x, y: minY }, coords) || !inPolygon({ x, y: maxY }, coords)) {
                    enclosed = false
                    break
                }
            }

            if (!enclosed) return

            if (area > maxArea) maxArea = area
        })
    })

    return maxArea
}

/**
 * @param {{x: number, y: number}} point
 * @param {{x: number, y: number}[]} coords
 * @returns {boolean}
 */
function inPolygon(point, coords) {
    const x = point.x
    const y = point.y

    // Check cache first
    const key = x + '.' + y
    if (pointsInPolygon.has(key)) return true
    if (pointsOutsidePolygon.has(key)) return false

    const res = isPointInPolygon(point, coords)

    if (res) pointsInPolygon.add(key)
    else pointsOutsidePolygon.add(key)

    return res


}

