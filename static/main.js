Vue.component('budget-display', {
    template: "#budget",
    props: ["balance", "daily", "remainingDays"]
});


Vue.component('transaction-input', {
    template: "#transaction-input",
    data: function () {
        return {
            description: undefined,
            amount: undefined
        }
    },
    methods: {
        add() {
            var self = this;
            axios.post('/api/transaction', {
                description: this.description,
                amount: this.amount
            }).then(function (response) {
                // TODO ML Update event and listener in parent
                // self.fetchBudget();
                // self.fetchTransactions();
            });
        },
        sub() {
            var self = this;
            axios.post('/api/transaction', {
                description: this.description,
                amount: "-" + this.amount
            }).then(function (response) {
                // self.fetchBudget();
                // self.fetchTransactions();
            });
        },
    }
});

var app = new Vue({
    el: '#app',
    data: {
        transactions: [],
        budget: {},

        // For entering a new transaction.
        amount: "",
        description: ""
    },
    created() {
        this.fetchTransactions();
        this.fetchBudget();
    },

    methods: {
        isIncome(index) {
            return this.transactions[index].amount > 0;
        },

        fetchTransactions() {
            axios.get('/api/transaction/2018/3').then(response => {
                this.transactions = response.data;
            });
        },

        fetchBudget() {
            axios.get('/api/transaction/2018/3/budget').then(response => {
                this.budget = response.data;
            });
        },

        getFormattedDate(date) {
            var year = date.getFullYear();
            var month = (1 + date.getMonth()).toString();
            month = month.length > 1 ? month : '0' + month;
            var day = date.getDate().toString();
            day = day.length > 1 ? day : '0' + day;
            return year + '-' + month + '-' + day;
        }


    }
})

