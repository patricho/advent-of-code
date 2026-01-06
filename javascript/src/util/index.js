import { readFileSync } from 'fs'

export const DIRECTIONS_AND_DIAGONALS = [
    { x: 0, y: -1 }, // up
    { x: 0, y: 1 }, // down
    { x: 1, y: 0 }, // right
    { x: -1, y: 0 }, // left
    { x: -1, y: -1 }, // up-left
    { x: -1, y: 1 }, // down-left
    { x: 1, y: -1 }, // up-right
    { x: 1, y: 1 }, // down-right
]

/**
 * @param {string[][]} grid
 * @param {number} y
 * @param {number} x
 * @returns {boolean}
 */
export function outOfBounds(grid, y, x) {
    return y < 0 || y >= grid.length || x < 0 || x >= grid[y].length
}

/**
 * @param {import("fs").PathOrFileDescriptor} filename
 * @returns {string[]}
 */
export function readFileLines(filename) {
    return readFileSync(filename, 'utf8').trim().split('\n')
}

export function readFileString(filename) {
    return readFileSync(filename, 'utf8').trim()
}

export function readFileNumberGrid(filename, separator = ',') {
    return readFileSync(filename, 'utf8')
        .trim()
        .split('\n')
        .map((line) => {
            return line.split(separator).map((s) => parseInt(s))
        })
}

/**
 * @param {string} label
 * @param {number} want
 * @param {() => number} got
 */
export function assert(label, want, got) {
    const start = performance.now()
    const res = got()
    const duration = performance.now() - start
    const match = want === res

    const icon = match ? '✓' : '✗'

    const white = '\x1b[97m'
    const lgray = '\x1b[37m'
    const gray = '\x1b[90m'
    const green = '\x1b[32m'
    const red = '\x1b[31m'
    const color = match ? green : red
    const reset = '\x1b[0m'

    console.log(
        `${color}${icon} ${lgray}${label} ${gray}want ${white}${want}${gray} got ${color}${res}${reset} ${gray}(${duration.toFixed(2)} ms)${reset}`
    )
}

export function prettyPrint(func) {
    const gray = '\x1b[90m'
    const reset = '\x1b[0m'

    console.log(`\n${gray}---${reset}\n`)

    func()
}

/**
 * @param {{x: number, y: number}} point
 * @param {{x: number, y: number}[]} coords
 * @returns {boolean}
 */
export function isPointInPolygon(point, coords) {
    const x = point.x
    const y = point.y

    let inside = false

    // Ray casting algorithm
    for (let i = 0, j = coords.length - 1; i < coords.length; j = i++) {
        const xi = coords[i].x
        const yi = coords[i].y
        const xj = coords[j].x
        const yj = coords[j].y

        // Check if the ray from the point crosses this edge
        const intersect = yi > y !== yj > y && x < ((xj - xi) * (y - yi)) / (yj - yi) + xi

        if (intersect) {
            inside = !inside
        }
    }

    if (inside) {
        return true
    }

    // Check if we're exactly on a border
    for (let i = 0, j = coords.length - 1; i < coords.length; j = i++) {
        if (isPointOnSegment(point, coords[i], coords[j])) {
            return true
        }
    }

    return false
}

/**
 * @param {{x: number, y: number}} point
 * @param {{x: number, y: number}} p1
 * @param {{x: number, y: number}} p2
 * @returns {boolean}
 */
export function isPointOnSegment(point, p1, p2) {
    // Check if point is within bounding box of segment
    const minX = Math.min(p1.x, p2.x)
    const maxX = Math.max(p1.x, p2.x)
    const minY = Math.min(p1.y, p2.y)
    const maxY = Math.max(p1.y, p2.y)

    if (point.x < minX || point.x > maxX || point.y < minY || point.y > maxY) {
        return false
    }

    // Use cross product to check if point is on the line
    // If cross product is 0, the point is collinear with the segment
    const crossProduct = (point.y - p1.y) * (p2.x - p1.x) - (point.x - p1.x) * (p2.y - p1.y)

    return Math.abs(crossProduct) < 1
}
