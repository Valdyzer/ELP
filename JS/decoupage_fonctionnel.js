const prompt = require("prompt-sync")({sigint:true}); //user input library

var checkword = require("check-if-word"), words = checkword('en'); //word check library

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

var players = (1,2)
var game_playing = true;
var i = 1;

var player_letters = [[],[]]

var player_grid = [[],[]]

function pick6Letters(letters) {
    //works

    var letter_list = []
    for (let i=0; i<6; i++){
        rdm = pickLetter(letters)
        letter_list.push(rdm)
    }
    return letter_list
}

function userInput(text,player) {
    //works
    console.log(text);
    var input_check = false;
    text = "Please enter a word: \n";
    while (input_check==false){
        console.log(text);
        var user_input = prompt();
        //console.log("User entered:", user_input);
        input_check = inputCheck(user_input,player);
    }
    let res = [user_input,input_check]
    return res;

}

function inputCheck(user_input,player) {
    //returns true if word passes check, false otherwise
    //3 letters min, inside the noms_communs dic (feminins, pluriels, verbes à tous les temps/modes/personnes),
    // exclus (prefixes, interjections, sigles, symboles, mots composés)
    //works

    user_input = user_input.toUpperCase();
    let split_word = user_input.split('');
    var res_string = ""
    
    if (split_word.length >= 3){

        for (let letter of split_word) {

            if (player.includes(letter)==false){
                console.log("Invalid Entry. Please only use letters you have.\n");
                return false
            }            
        }

        if (words.check(user_input.toLowerCase())) { 
            return true 
        } 
        else { 
            res_string+="Word does not exist...\n" 
        }
    }
    else { res_string += "Word not long enough, make sure the word is at least 3 letters long.\n" }
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

function passTurn() {
    if (i == 1) {i=2;} else {i=1;}
}

function exchange3Letters() {

}

function coup2Jarnac(){

}

function coup2Jarnac2() {
    console.log("Hello World!");
}

function turn(player) {
    console.log("\n\nIt's player ",player,"'s turn!\n");
    player -= 1;
    player_letters[player] = pick6Letters()
    console.log("You have 6 letters: ",player_letters[player]);
    while (true) {
        console.log("Do you want to make a word ? (yes/no)")
        choice = prompt().toLowerCase()
        if (choice == "yes") {
            let user_input = userInput("Make a word.",player_letters[player]);
            if (user_input[1]) { //if input is valid ==> add the word to grid and do rest of the play.
                player_grid[player].push(user_input[0]);
                console.log(player_grid[player]);
                break
            }
        }
        else if (choice == "no") {
            break  // PASSE SON TOUR (A coder)
        }
        else {
            console.log("\nINCORRECT !! Please answer by 'yes' or 'no'")
        }
    }
    

    
}



while (game_playing){
    turn(i);
    passTurn();
}



