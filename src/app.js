document.addEventListener('alpine:init', () => {

    Alpine.store('user', {
        id: null,
        username: null,
        email: null,
        isLogged: false,
        login(id, username, email) {
            this.id = id;
            this.username = username;
            this.email = email;
            this.isLogged = true;
        },
        logout() {
            this.id = null;
            this.username = null;
            this.email = null;
            this.isLogged = false;
        }
    });


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


    Alpine.data('orders', () => ({
        orders: [],
        loading: true,
        menus: [
            {id: 1, restaurantid: 1, name: 'Deserts', icon: 'fas fa-ice-cream', link: '#'},
            {id: 2, restaurantid: 1, name: 'Drinks', icon: 'fas fa-cocktail', link: '#'},
            {id: 3, restaurantid: 2, name: 'Burgers', icon: 'fas fa-hamburger', link: '#'},
            {id: 4, restaurantid: 2, name: 'Pizza', icon: 'fas fa-pizza-slice', link: '#'},
            {id: 5, restaurantid: 3, name: 'Sushi', icon: 'fas fa-fish', link: '#'},
            {id: 6, restaurantid: 3, name: 'Salads', icon: 'fas fa-carrot', link: '#'},
            {id: 7, restaurantid: 4, name: 'Pasta', icon: 'fas fa-pepper-hot', link: '#'},
            {id: 8, restaurantid: 4, name: 'Sandwiches', icon: 'fas fa-bread-slice', link: '#'},
        ],
        products: [
            {
                id: 1,
                menuid: 1,
                name: 'Chocolate Cake',
                description: 'Delicious chocolate cake',
                price: 10,
                image: 'https://source.unsplash.com/300x150/?ChocolateCake'
            },
            {
                id: 9,
                menuid: 1,
                name: 'Cheese Cake',
                description: 'Delicious cheese cake',
                price: 12,
                image: 'https://source.unsplash.com/300x150/?CheeseCake'
            },
            {
                id: 10,
                menuid: 1,
                name: 'Apple pie',
                description: 'Delicious apple pie',
                price: 12,
                image: 'https://source.unsplash.com/300x150/?ApplePie'
            },

            {
                id: 2,
                menuid: 2,
                name: 'Coca-Cola',
                description: 'Coca-Cola 33cl',
                price: 2,
                image: 'https://source.unsplash.com/300x150/?Coca-Cola'
            },
            {
                id: 11,
                menuid: 2,
                name: 'Fanta',
                description: 'Fanta 33cl',
                price: 2,
                image: 'https://source.unsplash.com/300x150/?Fanta'
            },
            {
                id: 12,
                menuid: 2,
                name: 'Sprite',
                description: 'Sprite 33cl',
                price: 2,
                image: 'https://source.unsplash.com/300x150/?Sprite'
            },

            {
                id: 3,
                menuid: 3,
                name: 'Cheese Burger',
                description: 'Cheese Burger with fries',
                price: 15,
                image: 'https://source.unsplash.com/300x150/?cheeseburger'
            },
            {
                id: 13,
                menuid: 3,
                name: 'Chicken Burger',
                description: 'Chicken Burger with fries',
                price: 15,
                image: 'https://source.unsplash.com/300x150/?chickenburger'
            },
            {
                id: 14,
                menuid: 3,
                name: 'Fish Burger',
                description: 'Fish Burger with fries',
                price: 15,
                image: 'https://source.unsplash.com/300x150/?fish burger'
            },

            {
                id: 4,
                menuid: 4,
                name: 'Pizza Margarita',
                description: 'Pizza Margarita 4 seasons',
                price: 20,
                image: 'https://source.unsplash.com/300x150/?pizza'
            },
            {
                id: 15,
                menuid: 4,
                name: 'Pizza 4 seasons',
                description: 'Pizza 4 seasons',
                price: 20,
                image: 'https://source.unsplash.com/300x150/?pizza'
            },
            {
                id: 16,
                menuid: 4,
                name: 'Pizza 4 cheese',
                description: 'Pizza 4 cheese',
                price: 20,
                image: 'https://source.unsplash.com/300x150/?pizza'
            },

            {
                id: 5,
                menuid: 5,
                name: 'Sushi Mix',
                description: 'Sushi Mix 24 pieces',
                price: 30,
                image: 'https://source.unsplash.com/300x150/?sushi'
            },
            {
                id: 17,
                menuid: 5,
                name: 'Sushi californian roll',
                description: 'Sushi Mix 24 pieces',
                price: 30,
                image: 'https://source.unsplash.com/300x150/?sushi'
            },
            {
                id: 18,
                menuid: 5,
                name: 'Sushi Mix 24 pieces',
                description: 'Sushi Mix 24 pieces',
                price: 30,
                image: 'https://source.unsplash.com/300x150/?sushi'
            },

            {
                id: 6,
                menuid: 6,
                name: 'Salad',
                description: 'Salad with vegetables',
                price: 8,
                image: 'https://source.unsplash.com/300x150/?salad'
            },
            {
                id: 19,
                menuid: 6,
                name: 'Salad with vegetables',
                description: 'Salad with vegetables',
                price: 8,
                image: 'https://source.unsplash.com/300x150/?salad'
            },

            {
                id: 7,
                menuid: 7,
                name: 'Pasta Carbonara',
                description: 'Pasta Carbonara with beef',
                price: 12,
                image: 'https://source.unsplash.com/300x150/?Pasta'
            },
            {
                id: 20,
                menuid: 7,
                name: 'Pasta with tomato sauce',
                description: 'Pasta with tomato sauce',
                price: 12,
                image: 'https://source.unsplash.com/300x150/?Pasta'
            },
            {
                id: 21,
                menuid: 7,
                name: 'Pasta with cheese',
                description: 'Pasta with cheese',
                price: 12,
                image: 'https://source.unsplash.com/300x150/?Pasta'
            },

            {
                id: 8,
                menuid: 8,
                name: 'Sandwich',
                description: 'Sandwich with chicken and cheese',
                price: 6,
                image: 'https://source.unsplash.com/300x150/?Sandwich'
            },
            {
                id: 22,
                menuid: 8,
                name: 'Sandwich with fish',
                description: 'Sandwich with fish',
                price: 6,
                image: 'https://source.unsplash.com/300x150/?Sandwich'
            },
        ]
    }))

    Alpine.store('cart', {
        items: [],
        total: 0,
        count: 0,
        add(newItem) {
            const cartItem = this.items.find(item => item.id === newItem.id);
            if (!cartItem) {
                this.items.push({...newItem, count: 1, total: newItem.price});
                this.total += newItem.price;
                this.count += 1;
            } else {
                this.items = this.items.map((item) => {
                    if (item.id !== newItem.id) return item;
                    item.count += 1;
                    item.total = item.price * item.count;
                    this.total += item.price;
                    this.count += 1;
                    return item;
                });
            }
            console.log("items :", this.items)
        },
        remove(id) {
            const cartItem = this.items.find(item => item.id === id);
            if (cartItem.count > 1) {
                this.items = this.items.map((item) => {
                    if (item.id !== id) return item;
                    item.count -= 1;
                    item.total = item.price * item.count;
                    this.total -= item.price;
                    this.count -= 1;
                    return item;
                });
            } else if (cartItem.count === 1) {
                this.items = this.items.filter((item) => item.id !== id);
                this.total -= cartItem.price;
                this.count -= 1;
            }
            console.log("items :", this.items)

        }
    });

    Alpine.data('cart_container', () => ({
        open: false,
        submitted: false,
        toggle() {
            if (this.submitted) {
                this.submitted = false;
            }

            this.open = !this.open;
            if (this.open) {
                document.body.style.overflow = 'hidden';
            } else {
                document.body.style.overflow = 'auto';
            }
        },
        submit() {

            console.log("submit :", JSON.stringify(Alpine.store('cart').items))

            const restaurantId = this.$el.getAttribute('data-restaurant-id');

            fetch(`http://localhost:8097/restaurant/${restaurantId}/create-order`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(Alpine.store('cart').items),
            })
                .then(response => (response.json()))
                .then(data => {
                    if (localStorage.getItem('orderData') !== null) {
                        let oldOrderData = JSON.parse(localStorage.getItem('orderData'));
                        localStorage.setItem('orderData', JSON.stringify([...oldOrderData, data]));
                    } else {
                        localStorage.setItem('orderData', JSON.stringify([data]));
                    }

                    this.toggle();
                    Alpine.store('cart').items = [];
                    Alpine.store('cart').total = 0;
                    Alpine.store('cart').count = 0;
                    this.submitted = true;
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
        }
    }));
})
