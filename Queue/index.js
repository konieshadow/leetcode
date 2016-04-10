/**
 * @constructor
 */
var Queue = function() {
    this._stack = [];
};

/**
 * @param {number} x
 * @returns {void}
 */
Queue.prototype.push = function(x) {
    this._stack.push(x);
};

/**
 * @returns {void}
 */
Queue.prototype.pop = function() {
    'use strict';
    let tempStack = [];
    while (this._stack.length > 1) {
        tempStack.push(this._stack.pop());
    }
    this._stack.pop();
    while(tempStack.length > 0) {
        this._stack.push(tempStack.pop());
    }
};

/**
 * @returns {number}
 */
Queue.prototype.peek = function() {
    'use strict';
    let tempStack = [];
    let top = null;
    while (this._stack.length > 0) {
        if (this._stack.length === 1) {
            top = this._stack.pop();
            tempStack.push(top);
        } else {
            tempStack.push(this._stack.pop());            
        }
    }
    while(tempStack.length > 0) {
        this._stack.push(tempStack.pop());
    }
    return top;
};

/**
 * @returns {boolean}
 */
Queue.prototype.empty = function() {
    return this._stack.length === 0;
};


/**
 * mytest
 */
var queue = new Queue();
queue.push(1);
queue.push(2);
queue.push(3);
queue.push(4);
queue.push(5);

console.log(queue);

queue.pop();
console.log(queue);

console.log(queue.peek());
console.log(queue);

console.log(queue.empty());