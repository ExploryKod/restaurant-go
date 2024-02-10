/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/templates/*.gohtml",
    "./src/templates/**/*.gohtml",
    "./src/*css",
    "./node_modules/flowbite/**/*.js"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

