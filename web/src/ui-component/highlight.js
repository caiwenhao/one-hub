import hljs from 'highlight.js/lib/core';
import json from 'highlight.js/lib/languages/json';
import python from 'highlight.js/lib/languages/python';
import javascript from 'highlight.js/lib/languages/javascript';
import objectivec from 'highlight.js/lib/languages/objectivec';
import bash from 'highlight.js/lib/languages/bash';
import shell from 'highlight.js/lib/languages/shell';

hljs.registerLanguage('json', json);
hljs.registerLanguage('python', python);
hljs.registerLanguage('javascript', javascript);
hljs.registerLanguage('js', javascript);
hljs.registerLanguage('node', javascript);
hljs.registerLanguage('nodejs', javascript);
hljs.registerLanguage('objectivec', objectivec);
hljs.registerLanguage('bash', bash);
hljs.registerLanguage('shell', shell);
hljs.registerLanguage('sh', bash);
hljs.registerLanguage('curl', bash);

export default hljs;
