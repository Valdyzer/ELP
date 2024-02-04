const prompt = require("prompt-sync")({sigint:true}); //user input library
const wordExists = require('word-exists');

var letter_dict = {
    A:14,
    B:4,
    C:7,
    D:5,
    E:19,
    F:2,
    G:4,
    H:2,
    I:11,
    J:1,
    K:1,
    L:6,
    M:5,
    N:9,
    O:8,
    P:4,
    Q:1,
    R:10,
    S:7,
    T:9,
    U:8,
    V:2,
    W:1,
    X:1,
    Y:1,
    Z:2
}

var players = (0,1)
var game_playing = true;
var first_turn_played = false;
var i = 1;

var player_letters = [[],[]]

var player_grid = [[],[]]


function pick6Letters() {
    //works

    var letter_list = []
    for (let i=0; i<6; i++){
        rdm = pickLetter()
        letter_list.push(rdm)
    }
    return letter_list
}

function userInput(text, player_array_letters, player) {
    //works
    var input_check = false;
    var input_exit = false;
    console.log("Here is your grid: ", player_grid[player]);
    text = "\nPlease enter a word. \n";
    
    console.log(text,"\nYou have ", player_letters[player].length, " letters: ", player_letters[player]);
    console.log("\nTo go back, type exit() and press enter.\n")
    var user_input = prompt();
    if (user_input=="exit()"){
        console.log("\nGoing back to menu...\n");
        input_exit = true;
    }
    else {
        input_check = inputCheck(user_input, player_array_letters, player);
    }
    if (input_exit==false && input_check==false){
        console.clear()
        return userInput(text, player_array_letters, player);
    }
    let res = [user_input,input_check];
    //console.log(res);
    return res;

}

function inputCheck(user_input, player_array_letters, player) {
    //returns true if word passes check, false otherwise
    //3 letters min, inside the noms_communs dic (feminins, pluriels, verbes à tous les temps/modes/personnes),
    // exclus (prefixes, interjections, sigles, symboles, mots composés)
    //works

    user_input = user_input.toUpperCase();
    let split_word = user_input.split('');
    var res_string = "";
    var tmp_player = player_array_letters.slice();
    
    if (split_word.length >= 3){

        for (let letter of split_word) {

            if (tmp_player.includes(letter)==false) {
                console.log("Invalid Entry. Please only use letters you have.\n");
                return false;
            }            
            const index = tmp_player.indexOf(letter);
            tmp_player.splice(index,1);
        }

        if (wordExists(user_input.toLowerCase()) == true) {
            return true;
        }
        else {res_string+="Word does not exist...\n";}
    }
    else{res_string += "Word not long enough, make sure the word is at least 3 letters long.\n";}
    console.log(res_string);
    return false;
}
    

function pickLetter() {
    //picks a random letter in the existing letters, reduces letter counter by one 
    //if enough letters left, deletes letter from dict otherwise
    //works

    var keys_array = Object.keys(letter_dict);
    var rdm_index = Math.floor(Math.random() * keys_array.length);
    var rdm_letter = keys_array[rdm_index];

    if (letter_dict[rdm_letter]==1){   
        delete letter_dict[rdm_letter];
    }
    else{letter_dict[rdm_letter] -= 1;}

    return rdm_letter
}

function whatToDo(player) {
    while (true) {
        console.clear();
        console.log("IT'S PLAYER ", player+1, "'S TURN")
        console.log("\nHere is your grid: ", player_grid[player])
        console.log("\nYou have ",player_letters[player].length," letters: ", player_letters[player]);
        console.log("\nWhat do you want to do ? (select a number)\n 1 - Write a new word \n 2 - Turn a word into a new one \n 3 - Pass")
        var choice = prompt()
        if (["1","2","3","4"].includes(choice)) {
            return choice
        }     
        else {console.log("\Invalid input. Try again.\n")}
    }
}

function newWord(player){
    console.clear();
    //console.log("Your grid: ", player_grid[player])
    let user_input = userInput("Make a word: ", player_letters[player], player);
    console.log(user_input);
    if (user_input[1]){ //if input is valid ==> add the word to grid and do rest of the play.
        player_grid[player].push(user_input[0]);
        split_word = user_input[0].split('');
        for (let letter of split_word) {
            console.log("Letters of split_word: ",letter);
            const index = player_letters[player].indexOf(letter.toUpperCase());
            if (index > -1){
                player_letters[player].splice(index,1);
                console.log("removed letter from array: ", player_letters[player])
            }
        }
        //console.log(player_grid[player]);
        return 1;
    }
    else if (user_input[0]=="exit()"){
        //console.log("\nReturning 0 because exit() was called\n");
        return -1;
    }
}

function chooseWord(player){//returns the index of the chosen word inside player_grid
    let user_input = "";
    let index_array = []
    for (let i = 1; i<= player_grid[player].length; i++){
        index_array.push(i.toString());
    }
    while (index_array.includes(user_input)==false){
        console.clear();
        console.log("What word do you want to use ?\n");
        for (let i = 0; i< index_array.length; i++){
            console.log((i+1), ".  ", player_grid[player][i],"\n");
        }
        console.log("Please choose a word by entering the corresponding number.\n");
        user_input = prompt();
    }
    return (parseInt(user_input)-1);
}

function changeWord(player){
    let old_word_index = chooseWord(player);
    var old_word = (player_grid[player][old_word_index].toUpperCase()).split('');
    console.log(old_word);
    console.log("\nPlease enter the new word to be added:\n");
    let new_word = prompt();
    let new_word_sliced = new_word.split('');
    let player_array = player_letters[player].concat(old_word); 
    let word_check = inputCheck(new_word, player_array, player);
    console.log(word_check);
    if (word_check==false){
        return word_check;
    }
    for (let letter of old_word){//removes the letters that have already been played from the array, only new letters left 
        if (new_word_sliced.includes(letter.toLowerCase())) {
            const index = new_word_sliced.indexOf(letter.toLowerCase());
            new_word_sliced.splice(index,1);
        }
    }
    // console.log(new_word_sliced);
    // console.log(player_letters[player]);
    for (let letter of new_word_sliced){
        const index = player_letters[player].indexOf(letter.toUpperCase());
        player_letters[player].splice(index,1);
    }
    // console.log(player_letters[player]);
    // console.log(new_word_sliced);
    // console.log(player_grid[player][old_word_index]);
    player_grid[player][old_word_index] = new_word;
    // console.log(player_grid[player][old_word_index]);
    return new_word;
}

function passTurn() {
    if (i == 0) {i=1;} else {i=0;}
}

function exchange3Letters() {

}

function writeJarnac(array, player_to_jarnac, player, grid, num) {
    console.clear()
    player_letters[player_to_jarnac] = array
    console.log("Write a word with these letters : ", player_letters[player_to_jarnac])
    var user_input = prompt().toLowerCase()
    split_word = user_input.split('');
    if (grid == "yes") {
        for (let letter of (player_grid[player_to_jarnac][num].toUpperCase())) {
            if (user_input.includes(letter) == false) {
                console.log("Invalid Entry. Please only use letters you have.\n");
                return false;    
            }
        }
    }
    var input_check = inputCheck(user_input, player_letters[player_to_jarnac], player_to_jarnac)
    if (input_check == true) {
        player_grid[player].push(user_input);
        for (letter of split_word) {
            // si le nouveau mot comporte toutes les lettres de l'ancien alors continue et return true sinon return false
            console.log("Letters of split_word: ", letter);
            const index = player_letters[player_to_jarnac].indexOf(letter.toUpperCase());
            if (index > -1) {
                player_letters[player_to_jarnac].splice(index,1);
                console.log("removed letter from array: ", player_letters[player_to_jarnac])
                }
        }
    }
    return true
}

function coupJarnac(player) {
    console.clear();
    console.log("IT'S PLAYER ", player+1, "'S TURN\n")
    player_to_jarnac = (player+1)%2
    console.log("Here is the player ", player_to_jarnac + 1, "'s grid : ", player_grid[player_to_jarnac])
    console.log("And here is the player ", player_to_jarnac + 1, "'s letters : ", player_letters[player_to_jarnac])
    console.log("\nDo you want to JARNAC ? (yes/no)")
    choice = prompt().toLowerCase()
    if (choice == "yes") {
        if (player_grid[player_to_jarnac].length != 0) {
            console.log("Do you need a word from the grid ? (yes/no)")
            grid = prompt().toLowerCase()
            if (grid == "yes") {
                const num = chooseWord(player_to_jarnac)
                word = (player_grid[player_to_jarnac][num].toUpperCase()).split('')
                array = player_letters[player_to_jarnac].concat(word)
                console.log(array)
                prompt()
                valid = writeJarnac(array, player_to_jarnac, player, grid, num)
                if (valid != true)
                    return coupJarnac(player)
                player_grid[player_to_jarnac].splice(num, 1)
            }
            else if (grid =="no") {
                valid = writeJarnac(player_letters[player_to_jarnac], player_to_jarnac, player, grid, 0)
                if (valid != true)
                    return coupJarnac(player)
            }  
        }
        else {
            valid = writeJarnac(player_letters[player_to_jarnac], player_to_jarnac, player, grid, 0)
            if (valid != true)
            return coupJarnac(player)
        }
    }   
    else if (choice == "no") 
        return 
    else
        return coupJarnac(player)
}


function menu(player) {
    console.clear();
    player_choice = whatToDo(player);

    if (player_choice == 1) { //if player chooses to place a new word
        let new_word_try = newWord(player);
        console.log(new_word_try)
        if (new_word_try==1){return turn(player);}
        else if (new_word_try==-1){
            return menu(player);
        }
    }
   else if (player_choice == 2) {
        console.log("User chose to change a word.");
        let change_word_try = changeWord(player);
        if (change_word_try==false){
            console.log("You have entered a non valid word. Please make sure the word exists and you have the required letters.\rPress enter to return to menu.")
            prompt();
            return menu(player);
        }
        console.log("New word selected: ",change_word_try);
        return turn(player);
        
   }
   else if (player_choice == 3) {
    console.log("User chose to skip turn");
    return 1;
   }
}

function firstTurn(player) {
    console.log("It's player ", player + 1, "'s turn!\n");
    player_letters[player] = pick6Letters()
    //console.log("You have ", player_letters[player].length," letters: ",player_letters[player]);
    return menu(player);

}

function turn(player) {
    while (true){
        console.log("It's player ",player + 1,"'s turn!\n");
        new_letter = pickLetter();
        player_letters[player].push(new_letter)
        return menu(player);
        }
}

while (game_playing) {
    if (first_turn_played == false) {
        firstTurn(i);
        passTurn();
        coupJarnac(i);
        firstTurn(i);
        passTurn();
        first_turn_played = true;
    }
    else {    
        coupJarnac(i);  
        turn(i);
        if (player_grid[i].length == 7) {
            // end_game
            break
        }
        else {passTurn()}
    }
}



