/** @type {import('tailwindcss').Config} */

module.exports = {
    content: ["./src/**/*.{html,ts}", "../internal/templates/*.html"], theme: {
        extend: {},
    }, plugins: [require("daisyui")], daisyui: {
        logs: false, themes: [{
            light: {
                ...require("daisyui/src/colors/themes")["[data-theme=light]"],
            }
        }]
    }
}