import PropTypes from 'prop-types';

import React, { useEffect, useRef, useState } from 'react';
import hljs from './highlight';
import { copy } from 'utils/common';

import 'assets/css/modern-code.css';

export default function CodeBlock({ language, code }) {
  const preRef = useRef(null);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    if (preRef.current && language) {
      // 清除之前的高亮
      preRef.current.removeAttribute('data-highlighted');
      preRef.current.className = `language-${language}`;
      hljs.highlightElement(preRef.current);
    }
  }, [code, language]);

  return (
    <div className="code-block" style={{ position: 'relative', margin: 0 }}>
      <pre className="hljs" style={{ margin: 0, padding: '1.5em', borderRadius: 0 }}>
        <code ref={preRef} className={`language-${language}`}>
          {code}
        </code>
      </pre>
      <button
        className="code-block__button"
        onClick={() => {
          copy(code);
          setCopied(true);
          setTimeout(() => {
            setCopied(false);
          }, 1500);
        }}
      >
        {copied ? '已复制' : '复制'}
      </button>
    </div>
  );
}

CodeBlock.propTypes = {
  language: PropTypes.string,
  code: PropTypes.string
};
