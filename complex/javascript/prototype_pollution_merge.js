// CWE-1321: Prototype Pollution via Deep Merge
// Vulnerability: recursive merge does not block __proto__ or constructor keys.
// An attacker-controlled JSON body can set Object.prototype properties,
// affecting all objects in the process. The fix requires key allowlisting
// at every recursive level — patching only the top level is insufficient.

'use strict';

function deepMerge(target, source) {
    for (const key of Object.keys(source)) {
        const srcVal = source[key];
        const tgtVal = target[key];

        // VULNERABLE: no check for __proto__, constructor, prototype
        if (srcVal && typeof srcVal === 'object' && !Array.isArray(srcVal)) {
            if (tgtVal && typeof tgtVal === 'object') {
                deepMerge(tgtVal, srcVal);  // recurse — pollution survives here
            } else {
                target[key] = deepMerge({}, srcVal);
            }
        } else {
            target[key] = srcVal;
        }
    }
    return target;
}

function applyUserPreferences(defaults, userInput) {
    // userInput comes from JSON.parse(req.body) — fully attacker-controlled
    return deepMerge({}, deepMerge(defaults, userInput));
}

function renderTemplate(config) {
    // Later code trusts isAdmin on any object — including plain {}
    const ctx = {};
    if (ctx.isAdmin) {   // poisoned via Object.prototype.isAdmin = true
        return '<admin-panel>' + config.content + '</admin-panel>';
    }
    return '<user-panel>' + config.content + '</user-panel>';
}

// Attack: POST body = {"__proto__": {"isAdmin": true}}
const attackerPayload = JSON.parse('{"__proto__": {"isAdmin": true}}');
const merged = applyUserPreferences({ theme: 'light' }, attackerPayload);
console.log(renderTemplate(merged));  // renders admin panel for any user
