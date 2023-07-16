/** @type {import("tailwindcss").Config} */
module.exports = {
    content: ["./src/**/*.{html,js,ts,tsx}", "./node_modules/tw-elements/dist/js/**/*.js"],
    theme: {
        extend: {
            colors: {
                primary: "#658fb5",
                "primary-darker": "#2a5780",
                card: {
                    desc: "#8c989c",
                    error: "#D63650"
                },
                main: {
                    "blue-1": "#497d9e",
                    "blue-2": "#4377bf",
                    "blue-3": "#c7d1d9",
                    "error": "#D63650", 
                    "red-1": "#fff2f0",
                    "red-2": "#ffccc7"
                }
            },
            maxWidth: {
                "xl-desc-size": "300px",
                "sm-desc-size": "209px"
            },
            minHeight: {
                "mobile-card": "600px"
            },
            height: {
                "mobile-card-height": "calc(100vh - 3.75rem)"
            }
        },
        screens: {
            "xs": {"max": "390px"},
            "sm": {"max": "639px"},
            "md": {"min": "390px", "max": "639px"},
            "lg": "640px"
        }
    },
    plugins: [require("tw-elements/dist/plugin")]
};
