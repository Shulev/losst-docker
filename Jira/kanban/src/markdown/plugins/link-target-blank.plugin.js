/**
 * Markdown-it extension
 *
 * Replaces default link_open renderer to add target=_blank
 * attribute to every link generated by markdown-it
 */
(function() {
    window.md_link_target_blank_plugin = function(md) {
        var defaultRenderer = md.renderer.rules.link_open || function(token, idx, options, env, self) {
            return self.renderToken(token, idx, options);
        };

        function linkOpenRenderer(tokens, idx, options, env, self) {
            // If you are sure other plugins can't add `target` - drop check below
            var aIndex = tokens[idx].attrIndex('target');
            var local = tokens[idx].attrIndex('data-link-local');

            if (local < 0){
                if (aIndex < 0) {
                    tokens[idx].attrPush(['target', '_blank']); // add new attribute
                } else {
                    tokens[idx].attrs[aIndex][1] = '_blank';    // replace value of existing attr
                }
            }

            // pass token to default renderer.
            return defaultRenderer(tokens, idx, options, env, self);
        }

        md.renderer.rules.link_open = linkOpenRenderer;
    };
}());