'use strict';

function myHandler() {
        throw new Error('My Error');
}

function main() {
        myHandler();
}

main();
