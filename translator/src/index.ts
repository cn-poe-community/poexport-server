import bodyParser from 'body-parser';
import { TranslatorFactory } from 'cn-poe-translator';
import { transform } from "pob-building-creater";
import express from 'express'
import { deflate } from "pako";
import { Base64 } from 'js-base64';

type TranslateRequest = {
    "items": unknown,
    "passiveSkills": unknown
}

const factory = TranslatorFactory.Default();
const jsonTranslator = factory.getJsonTranslator();

const app = express()
app.use(bodyParser.json({ limit: '5mb' }))

app.post('/pob/create', (req, res) => {
    const r = req.body as TranslateRequest
    jsonTranslator.translateItems(r.items)
    jsonTranslator.translatePassiveSkills(r.passiveSkills)

    const building = transform(r.items, r.passiveSkills)

    const compressed = deflate(building.toString())
    let code = Base64.fromUint8Array(compressed);
    res.send(code.replaceAll("+", "-").replaceAll("/", "_"))
})

process.on('uncaughtException', err => {
    console.log('caught uncaught exception: ' + err)
})

const port = 8000
app.listen(port, () => {
    console.log(`Listening on port ${port}`)
})
