// https://github.com/delventhalz/advent-of-code-2023/blob/main/09-mirage-maintenance/1.js

const fs = require("fs");

const findDifferences = (history) => {
    const diffs = [];

    for (let i = 1; i < history.length; i += 1) {
        diffs.push(history[i] - history[i - 1]);
    }

    return diffs;
};

const findNextValue = (history) => {
    // If every value is the same, we know the next value in the sequence
    if (history.every((entry) => entry === history[0])) {
        return history[0];
    }

    // If not, we can get the next value by adding the next diff
    const diffs = findDifferences(history);
    const nextDiff = findNextValue(diffs);
    return last(history) + nextDiff;
};

const sum = (nums) => {
    return nums.reduce((prev, curr) => prev + curr, 0);
};

const last = (nums) => {
    return nums[nums.length - 1];
};

const main = (lines) => {
    const histories = lines.map((line) => line.split(" ").map(Number));
    const nextValues = histories.map(findNextValue);

    const res = sum(nextValues);

    nextValues.map((v) => console.log(v));
    console.log("result:", res);
};

const lines = fs.readFileSync("input.txt", "utf8");
console.log(lines);
main(lines.split(`\n`));
