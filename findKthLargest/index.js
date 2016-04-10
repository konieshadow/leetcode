/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findKthLargest = function(nums, k) {
    'use strict';
    let sortedNums = nums.sort((left, right) => right - left);
    return sortedNums[k-1];
};


/**
 * mytest
 */
var nums = [3,2,3,1,2,4,5,5,6,7,7,8,2,3,1,1,1,10,11,5,6,2,4,7,8,5,6];
var result = findKthLargest(nums, 2);
console.log(result);