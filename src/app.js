document.addEventListener('alpine:init', () => {

    Alpine.data('checker', () => ({
        username: '',
        email: '',
        error: {
            username: null,
            email: null
        },
        success: {
            username: null,
            email: null
        },
        checkUsername(value, type) {
            if (type === 'email' && !value.match(/^[^\s@]+@[^\s@]+\.[^\s@]+$/)) {
                this.error.email = 'Email is not valid'
                this.success.email = false
                return
            }
            fetch('http://localhost:8097/checkEmailAndUsername?username=' + (type === 'username' ? value : '') + '&email=' + (type === 'email' ? value : ''))
                .then((response) => {
                    if (!response.ok) {
                        throw new Error('Erreur lors de la requête.');
                    }
                    return response.json()
                })
                .then((json) => {

                    this.error.username = json.username?.exists ? json.username.message : null;
                    this.success.username = !json.username?.exists ? json.username.message : null;

                    this.error.email = json.email?.exists ? json.email.message : null;
                    this.success.email = !json.email?.exists ? json.email.message : null;

                })
                .catch((error) => {
                    console.log('Error during checking:', error)
                })
        }
    }))

    Alpine.data('stepper', () => ({
        selected: 1,
        total: 2,
        previous() {
            this.selected = Math.max(this.selected - 1, 1)
        },
        next() {
            this.selected = Math.min(this.selected + 1, this.total)
        }
    }))
})
