// Self-invoking anonymous function to encapsulate the extension and avoid polluting the global namespace
(function(){
    // Importing the HTMX internal API type for use within this extension
    /** @type {import("../htmx").HtmxInternalApi} */
    var api;

    // Prefix used to identify HTMX response target attributes in the DOM
    var attrPrefix = 'hx-target-';

    // Function to check if a string starts with a given prefix
    // Necessary for compatibility with browsers that do not support String.prototype.startsWith (e.g., IE11)
    function startsWith(str, prefix) {
        return str.substring(0, prefix.length) === prefix;
    }

    /**
     * Finds the target element for a given response code.
     * @param {HTMLElement} elt - The element to start the search from.
     * @param {number} respCodeNumber - The HTTP response code to find a target for.
     * @returns {HTMLElement | null} - The target element or null if not found.
     */
    function getRespCodeTarget(elt, respCodeNumber) {
        if (!elt || !respCodeNumber) return null;

        var respCode = respCodeNumber.toString();

        // '*' is the original syntax, as the obvious character for a wildcard.
        // The 'x' alternative was added for maximum compatibility with HTML
        // templating engines, due to ambiguity around which characters are
        // supported in HTML attributes.
        //
        // Start with the most specific possible attribute and generalize from
        // there.
        var attrPossibilities = [
            respCode,

            respCode.substr(0, 2) + '*',
            respCode.substr(0, 2) + 'x',

            respCode.substr(0, 1) + '*',
            respCode.substr(0, 1) + 'x',
            respCode.substr(0, 1) + '**',
            respCode.substr(0, 1) + 'xx',

            '*',
            'x',
            '***',
            'xxx',
        ];
        if (startsWith(respCode, '4') || startsWith(respCode, '5')) {
            attrPossibilities.push('error');
        }

        // Iterates through the possible attributes to find a target element 
        // based on the response code
        for (var i = 0; i < attrPossibilities.length; i++) {
            var attr = attrPrefix + attrPossibilities[i];
            var attrValue = api.getClosestAttributeValue(elt, attr);
            // Special handling for "this" value to find and return the current element
            if (attrValue) {
                if (attrValue === "this") {
                    return api.findThisElement(elt, attr);
                } else {
                    return api.querySelectorExt(elt, attrValue);
                }
            }
        }
        
        return null; // Return null if no matching target is found
    }

    // Handles error flags based on HTMX configuration settings
    /** @param {Event} evt */
    function handleErrorFlag(evt) {
        // Adjusts the isError flag on the event detail based on configuration settings
        if (evt.detail.isError) {
            if (htmx.config.responseTargetUnsetsError) {
                evt.detail.isError = false;
            }
        } else if (htmx.config.responseTargetSetsError) {
            evt.detail.isError = true;
        }
    }

    // Define the 'response-targets' extension
    htmx.defineExtension('response-targets', {

        // Initializes the extension with the HTMX API reference and configures default settings
        /** @param {import("../htmx").HtmxInternalApi} apiRef */
        init: function (apiRef) {
            api = apiRef;
            // Default configuration settings
            if (htmx.config.responseTargetUnsetsError === undefined) {
                htmx.config.responseTargetUnsetsError = true;
            }
            if (htmx.config.responseTargetSetsError === undefined) {
                htmx.config.responseTargetSetsError = false;
            }
            if (htmx.config.responseTargetPrefersExisting === undefined) {
                htmx.config.responseTargetPrefersExisting = false;
            }
            if (htmx.config.responseTargetPrefersRetargetHeader === undefined) {
                htmx.config.responseTargetPrefersRetargetHeader = true;
            }
        },

        /**
         * @param {string} name
         * @param {Event} evt
         */
        // Event handler for the extension that modifies behavior based on HTTP response status
        onEvent: function (name, evt) {
            // Custom logic for handling different HTTP status codes and determining if content should be swapped
            if (name === "htmx:beforeSwap"    &&
                evt.detail.xhr                &&
                evt.detail.xhr.status !== 200) {
                if (evt.detail.target) {
                    if (htmx.config.responseTargetPrefersExisting) {
                        evt.detail.shouldSwap = true;
                        handleErrorFlag(evt);
                        return true;
                    }
                    if (htmx.config.responseTargetPrefersRetargetHeader &&
                        evt.detail.xhr.getAllResponseHeaders().match(/HX-Retarget:/i)) {
                        evt.detail.shouldSwap = true;
                        handleErrorFlag(evt);
                        return true;
                    }
                }
                if (!evt.detail.requestConfig) {
                    return true;
                }
                var target = getRespCodeTarget(evt.detail.requestConfig.elt, evt.detail.xhr.status);
                if (target) {
                    handleErrorFlag(evt);
                    evt.detail.shouldSwap = true;
                    evt.detail.target = target;
                }
                return true;
            }
        }
    });
})();

/*
This code provides enhanced control over how HTMX handles server responses, particularly error 
statuses, by dynamically selecting target elements based on response codes and possibly changing 
the error handling behavior according to the configuration. It's a powerful way to add more 
nuanced client-server interaction patterns to your HTMX-enabled web application.
*/