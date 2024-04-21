/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./view/templates/**/*.{html,js}"],
    theme: {
        extend: {},
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
        require("daisyui"),
    ],

    daisyui: {
        themes: ["retro"],
    }
}
