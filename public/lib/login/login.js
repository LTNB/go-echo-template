// var bcrypt = dcodeIO.bcrypt;
// var salt = bcrypt.genSaltSync(12);

$(document).ready(() => {
    new Login('[data-id="email"]', '[data-id="password-interface"]', '[data-id="password"]');
});

class Login {
    constructor(email, passwordInterface, password) {
        this.email = $(email);
        this.passwordInterface = $(passwordInterface);
        this.password = $(password);
        // this.loadEventListener();
    }

    // loadEventListener() {
    //     this.passwordInterface.change((event) => {
    //         console.log("password: " + this.passwordInterface.val());
    //         let hashValue = bcrypt.hashSync(this.passwordInterface.val(), salt);
    //         console.log("hash password: " +hashValue)
    //         this.password.val(hashValue)
    //     })
    // }
}