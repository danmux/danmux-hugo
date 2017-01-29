var main = new Terminator(document.getElementById('contentWrapper'), 
            { 
                prefix: '<span class="green">guest@mullineux</span>:<span class="red">~$</span> ',
                alwaysFocus: true,
                autoScroll: (window.innerWidth >= 600)
            });

window.onresize = function() {
    console.log("Resizing!");
    if (window.innerWidth >= 600) {
        main.config.autoScroll = true;
    } else {
        main.config.autoScroll = false;
    }
};
            
function printClass(term, className) {
    var cloneTarget = null;
    if (className && (cloneTarget = document.getElementsByClassName(className))) {
        if (cloneTarget[0]) {
            term.writeHTML(cloneTarget[0].innerHTML);
        } else {
            term.writeLine("cat: " + className + ": No such file or directory");
        }
    }
}

main.register(function(term, command) {
    term.writeLine('welcome.txt');
    term.writeLine('contact.txt');
    term.writeLine('credits.txt');
    term.prompt();
}, 'ls');
			
main.register(function(term, command) {
    command = command.split(' ');
    var arg = command[1] || '';
    if (arg.indexOf('.txt') !== -1
        && arg.indexOf('.txt') === (arg.length - 4))
        arg = arg.substring(0, arg.indexOf('.txt'));
        
    printClass(term, arg);
    term.prompt();
    return;
}, 'cat');

main.register(function(term, command) {
	printClass(term, 'dan');
	term.prompt();
	return;
}, 'dan');

main.register(function(term, command) {
	printClass(term, 'fergus');
	term.prompt();
	return;
}, 'fergus');

main.register(function(term, command) {
	printClass(term, 'help');
	term.prompt();
	return;
}, ['help', 'man', '?']);

main.register(function(term, command) {
	term.writeLine('/users/guest');
	term.prompt();
}, 'pwd');

main.register(function(term, command) {
	term.element.innerHTML = '';
	term.prompt();
}, 'clear');

main.register(function(term, command) {
    term.writeLine('This user is not in the cders file. This incident has been reported.');
    term.prompt();
}, 'cd');

main.register(function(term, command) {
    term.writeLine('guest');
    term.prompt();
}, 'whoami');

main.register(function(term, command) {
    term.writeLine(command.split(' ').slice(1).join(' '));
    term.prompt();
}, 'echo');


main.register(function(term, command) {
    var contentWrapper = document.getElementById('contentWrapper');
    contentWrapper.classList.toggle('hacker');
    term.writeLine('Hacker mode: ' + (contentWrapper.classList.contains('hacker') ? 'ENABLED' : 'DISABLED'));
    term.prompt();
}, ['hack', 'hacker']);

main.prompt();
main.autoType('cat welcome.txt', 1000);