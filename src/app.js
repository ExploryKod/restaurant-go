document.addEventListener('alpine:init', () => {
const serverUrl = location.hostname;

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
            fetch(`${serverUrl}/checkEmailAndUsername?username=` + (type === 'username' ? value : '') + '&email=' + (type === 'email' ? value : ''))
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


    // Fonction helper pour obtenir une image Unsplash par mot-clé
    async function getUnsplashImage(query) {
        try {
            const protocol = window.location.protocol;
            const host = window.location.host;
            const port = window.location.port ? `:${window.location.port}` : '';
            const url = `${protocol}//${host}${port}/api/unsplash/image?query=${encodeURIComponent(query)}`;
            console.log('Fetching image for:', query, 'from:', url);
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error(`Failed to fetch image: ${response.status} ${response.statusText}`);
            }
            const data = await response.json();
            console.log('Image received for', query, ':', data.url);
            return data.url || `https://source.unsplash.com/300x150/?${encodeURIComponent(query)}`;
        } catch (error) {
            console.error('Error fetching Unsplash image for', query, ':', error);
            // Fallback vers l'ancienne URL
            return `https://source.unsplash.com/300x150/?${encodeURIComponent(query)}`;
        }
    }

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
            {id: 1, menuid: 1, name: 'Chocolate Cake', description: 'Delicious chocolate cake', price: 10, imageQuery: 'ChocolateCake', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 9, menuid: 1, name: 'Cheese Cake', description: 'Delicious cheese cake', price: 12, imageQuery: 'CheeseCake', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 10, menuid: 1, name: 'Apple pie', description: 'Delicious apple pie', price: 12, imageQuery: 'ApplePie', image: '/src/assets/hero/hero_plat_3.jpg'},
            {id: 2, menuid: 2, name: 'Coca-Cola', description: 'Coca-Cola 33cl', price: 2, imageQuery: 'Coca-Cola', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 11, menuid: 2, name: 'Fanta', description: 'Fanta 33cl', price: 2, imageQuery: 'Fanta', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 12, menuid: 2, name: 'Sprite', description: 'Sprite 33cl', price: 2, imageQuery: 'Sprite', image: '/src/assets/hero/hero_plat_3.jpg'},
            {id: 3, menuid: 3, name: 'Cheese Burger', description: 'Cheese Burger with fries', price: 15, imageQuery: 'cheeseburger', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 13, menuid: 3, name: 'Chicken Burger', description: 'Chicken Burger with fries', price: 15, imageQuery: 'chickenburger', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 14, menuid: 3, name: 'Fish Burger', description: 'Fish Burger with fries', price: 15, imageQuery: 'fish burger', image: '/src/assets/hero/hero_plat_3.jpg'},
            {id: 4, menuid: 4, name: 'Pizza Margarita', description: 'Pizza Margarita 4 seasons', price: 20, imageQuery: 'pizza', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 15, menuid: 4, name: 'Pizza 4 seasons', description: 'Pizza 4 seasons', price: 20, imageQuery: 'pizza', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 16, menuid: 4, name: 'Pizza 4 cheese', description: 'Pizza 4 cheese', price: 20, imageQuery: 'pizza', image: '/src/assets/hero/hero_plat_3.jpg'},
            {id: 5, menuid: 5, name: 'Sushi Mix', description: 'Sushi Mix 24 pieces', price: 30, imageQuery: 'sushi', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 17, menuid: 5, name: 'Sushi californian roll', description: 'Sushi Mix 24 pieces', price: 30, imageQuery: 'sushi', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 18, menuid: 5, name: 'Sushi Mix 24 pieces', description: 'Sushi Mix 24 pieces', price: 30, imageQuery: 'sushi', image: '/src/assets/hero/hero_plat_3.jpg'},
            {id: 6, menuid: 6, name: 'Salad', description: 'Salad with vegetables', price: 8, imageQuery: 'salad', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 19, menuid: 6, name: 'Salad with vegetables', description: 'Salad with vegetables', price: 8, imageQuery: 'salad', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 7, menuid: 7, name: 'Pasta Carbonara', description: 'Pasta Carbonara with beef', price: 12, imageQuery: 'Pasta', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 20, menuid: 7, name: 'Pasta with tomato sauce', description: 'Pasta with tomato sauce', price: 12, imageQuery: 'Pasta', image: '/src/assets/hero/hero_plats_1.jpg'},
            {id: 21, menuid: 7, name: 'Pasta with cheese', description: 'Pasta with cheese', price: 12, imageQuery: 'Pasta', image: '/src/assets/hero/hero_plat_3.jpg'},
            {id: 8, menuid: 8, name: 'Sandwich', description: 'Sandwich with chicken and cheese', price: 6, imageQuery: 'Sandwich', image: '/src/assets/hero/hero_plats.jpg'},
            {id: 22, menuid: 8, name: 'Sandwich with fish', description: 'Sandwich with fish', price: 6, imageQuery: 'Sandwich', image: '/src/assets/hero/hero_plats_1.jpg'},
        ],
        async init() {
            // Log pour déboguer
            console.log('Orders init - products count:', this.products.length);
            console.log('First product:', this.products[0]);
            
            // Optionnel : améliorer les images en arrière-plan via l'API Unsplash
            console.log('Improving images via Unsplash API for', this.products.length, 'products');
            this.products.forEach(async (product, index) => {
                try {
                    const imageUrl = await getUnsplashImage(product.imageQuery);
                    if (imageUrl && imageUrl !== product.image) {
                        product.image = imageUrl;
                        console.log(`Image improved for ${product.name}`);
                    }
                } catch (error) {
                    console.error(`Error improving image for ${product.name}:`, error);
                }
            });
        }
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

            fetch(`${serverUrl}/restaurant/${restaurantId}/create-order`, {
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

    // Alpine.data('time', () => (// Parse the timestamp using Moment.js
    //     {
    //         parsedTime: "",
    //         parser(timeStamp) {
    //             // Format the timestamp as HH:mm:ss
    //             this.parseTime = moment(timeStamp).format("HH:mm:ss")
    //         }
    //     }
    // ))
})
