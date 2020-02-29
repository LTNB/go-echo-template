$(document).ready(() => {
    new Layout();
});

class Layout {
    constructor() {
        this.selectLang = $('[data-id="select-lang"]')
        this.inputLang = $('[data-id="input-lang"]')
        this.loadEvent()
    }

    loadEvent(){
        this.selectLang.on('change', (e) =>{
            console.log(e.target.value)
        })
    }
}