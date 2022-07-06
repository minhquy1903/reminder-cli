const express = require("express")
const bodyParser = require("body-parser")
const notifier = require("node-notifier")
const path = require("path")

const app = express()
const port = process.env.PORT || 9000

app.use(bodyParser.json())

app.get("/heath" ,(req, res) => {
    res.status(200).send("Notifier is alive =))")
})

app.post("/notify", (req, res) => {
    notify(req.body, reply => res.send(reply))
})

app.listen(port, () => console.log(`Notifier is running in port ${port}`))

const notify = ({title, message}, cb) => {

    notifier.notify(
        {
            title: title || "Non title",
            message: message || "Non message",
            icon: path.join(__dirname, "honey-badger.png"),
            sound: true,
            wait: true,
            reply: true,
            closeLabel: "Ok",
            timeout: 8 
        },
        (err, response, reply) => cb(reply)
    )
}