
import inquirer from 'inquirer';
import exec from 'child_process';
import EventEmitter from 'events';
import { spawn } from 'child_process';




const emitter = new EventEmitter();
emitter.setMaxListeners(30)




/*	-----------	FONCTION MAIN RÉALISANT DES OPÉRATIONS ASYNCHRONE	-----------	*/

async function main() {

	var input = await inquirer.prompt(Read());
	var commandName = input.commande;

	if (commandName === 'lp') {
		await commandLP();
	}
	else if (commandName.indexOf('bing') !== -1) {
		await commandBing(commandName);
	}
	else if (commandName.indexOf('!') !== -1) {
		await commandBackground(commandName);
	}
	else if (commandName.indexOf('keep') !== -1) {		// A FINIR !!
		await commandKeep(commandName);
	}
	else if (commandName === '') {
		return main();
	}
	else {
		await commandExec(commandName);
	}

}




/*	-----------	PATH ACTUEL + COMMANDE À TAPER	-----------	*/

function Read()	{
	var currentDir = process.cwd();
	return  {	message: currentDir.replace(' ', '-') + ' %',
				name:'commande'	}
}




/*	-----------	LISTE DES PROCESSUS EN COURS (NUM + ID + NOM DE LA COMMANDE)	-----------	*/

function commandLP() {
	exec('ps', (error, stdout, stderr) => { 
    	if (error) {
        	console.error('\n' + error);
      	}
      	else if (stderr) {
        	console.log('stderr : ' + stderr);
      	}
      	var tab = stdout.split('\n');		//	tableau dont chaque élément correspond à un processus en cours
      	tab = tab.slice(0, tab.length - 1);

      	for (let i=0 ; i<tab.length ; i++) {
        	if (i === 0) {
        		console.log('\t' + tab[i]);
        	}
        	else {
        		console.log(i + '.\t' + tab[i]);
        	}
      	}
      	console.log('\n');
      	return main();
  	});
}




/* -----------	LISTE DES PROCESSUS EN COURS (NUM + ID + NOM DE LA COMMANDE)	-----------	*/

function commandBing(cmd) {
	var tab = cmd.split(' ');
	if (tab.length !== 3) {
		return commandExec(cmd);
	}
	switch(tab[1]) {
		case '-k':
			commandExec('kill ' + tab[2]);				// tue
			break;
		case '-p':
			commandExec('kill -STOP ' + tab[2]);		// pause
			break;
		case '-c':
			commandExec('kill -CONT ' + tab[2]);		// reprend
			break;
		default:
			commandExec(cmd);
			break;
	}
}




/*	-----------	EXÉCUTION D'UN PROGRAMME EN TACHE DE FOND	-----------	*/

function commandBackground(cmd) {							// ne marche pas même en utilisant directement la commande avec '&''
	var tab = cmd.split(' ');
	var replace = '';
	for (var i = 0; i < tab.length-1; i++) {
		replace = replace + tab[i] + ' ';
	}
	return commandExec(replace + '&');
}




/*	-----------	EXÉCUTION D'UNE COMMANDE QUELCONQUE	-----------	*/

function commandKeep(cmd){
	// A FINIR !!
  return main();
}




/*	-----------	EXÉCUTION D'UNE COMMANDE QUELCONQUE	-----------	*/

function commandExec(cmd) {
	exec(cmd, (error, stdout, stderr) => {
  		if (error) {
    		console.error('\n' + error);
    	}
    	else if (stderr) {
        	console.log('stderr : ' + stderr);
        }
  		console.log(stdout);
  		
  		return main();
	});
}




/*	-----------	 ENTÊTE DU CLI	-----------	*/

console.log('\n\n\nWELCOME TO THIS SHELL (Ctrl + p to quit)');
console.log('Platform : ' + process.platform + '\n' );



main();
// quitter avec (Crtl + p)
process.stdin.setEncoding('utf8');
process.stdin.on('data', (key) => {
    if(key === '\u0010') { 
      	process.exit(); 
    }
});
