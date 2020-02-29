$(document).ready(() => {
    new Home('[data-id="body"]', '[data-id="flash-msg"]', '[data-id="data-table"]', '[data-id="btn-delete"]');
});

class Home {
    constructor(body, flashMsgElem, table, btnDeleteUser) {
        this.flashMsgElem = $(flashMsgElem);
        this.body = $(body);
        this.table = $(table);
        this.btnDeleteUser = $(btnDeleteUser);
        this.init()
        this.loadEventListener()
    }

    init() {
        this.table.DataTable({
            "paging": true,
            "lengthChange": false,
            "searching": false,
            "ordering": true,
            "info": true,
            "autoWidth": false,
        });
    }

    loadEventListener() {
        let self = this
        this.btnDeleteUser.on('click', function (e) {
            let id = $(this).parent('td').find('input:hidden').val();
            $.ajax({
                url: `/api/user/${id}`,
                type: 'DELETE',
                success: function (res) {
                    let msg = ""
                    if (res.data) {
                        $(`#${id}`).remove();
                    }
                    if (!$('[data-id="flash-msg"]')[0]) {
                        self.body.prepend(`<p class="alert alert-info" role="alert" data-id="flash-msg">${res.msg}</p>`);
                        self.flashMsgElem = $('[data-id="flash-msg"]')
                    } else {
                        // self.flashMsgElem.cleanData()
                        self.flashMsgElem.text(res.msg)
                    }
                }
            });
        })
    }

}
