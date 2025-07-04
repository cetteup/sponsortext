import { parseSponsorTextVariables } from "../lib/parse"

describe("parseSponsorTextVariables", () => {
    test("parses variables with suffix", () => {
        // GIVEN
        const text = "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "discord": "https://discord.gg/vx4AKRfj",
            "provider": "bf2hub.com",
        })
    })

    test("parses variables without suffix", () => {
        // GIVEN
        const text = "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "discord": "https://discord.gg/vx4AKRfj",
            "provider": "bf2hub.com",
        })
    })

    test("skips non-variables prefix", () => {
        // GIVEN
        const text = "Join our event this Sunday! $vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "discord": "https://discord.gg/vx4AKRfj",
            "provider": "bf2hub.com",
        })
    })

    test("stops before non-variables suffix", () => {
        // GIVEN
        const text = "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$ Apply to become an admin today!"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "discord": "https://discord.gg/vx4AKRfj",
            "provider": "bf2hub.com",
        })
    })

    test("reads escaped syntax characters", () => {
        // GIVEN
        const text = "$vars:\\$trange\\=key=https://example.com?query\\=\\$start\\;end$"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "$trange=key": "https://example.com?query=$start;end",
        })
    })

    test("reads escaped whitespaces", () => {
        // GIVEN
        const text = "$vars:four-spaces=\\ \\ \\ \\ $"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "four-spaces": "    ",
        })
    })

    test("omits irrelevant whitespaces", () => {
        // GIVEN
        const text = "$vars: trimmed = one  two  three $"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "trimmed": "one two three",
        })
    })

    test("ignores incomplete key", () => {
        // GIVEN
        const text = "$vars:discord=https://discord.gg/vx4AKRfj;website$"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "discord": "https://discord.gg/vx4AKRfj",
        })
    })

    test("includes key with no value", () => {
        // GIVEN
        const text = "$vars:discord=https://discord.gg/vx4AKRfj;teamspeak=$"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({
            "discord": "https://discord.gg/vx4AKRfj",
            "teamspeak": "",
        })
    })

    test("returns empty object for empty variables section", () => {
        // GIVEN
        const text = "$vars:$"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({})
    })

    test("returns empty object for sponsor text without variables", () => {
        // GIVEN
        const text = "Our server is the best!"

        // WHEN
        const variables = parseSponsorTextVariables(text)

        expect(variables).toEqual({})
    })
})
