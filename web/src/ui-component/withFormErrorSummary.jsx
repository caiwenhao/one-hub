// 高阶封装：提供表单错误汇总与定位
import { useRef } from 'react';
import FormErrorSummary from './FormErrorSummary';

export default function withFormErrorSummary(Component, { idPrefix = 'field-' } = {}) {
  return function Wrapped(props) {
    const refMap = useRef({});
    const bindRef = (name) => (node) => {
      if (node) refMap.current[name] = node;
    };
    const onJump = (name) => {
      const node = refMap.current[name];
      if (node && node.scrollIntoView) node.scrollIntoView({ behavior: 'smooth', block: 'center' });
      if (node && node.focus) node.focus();
    };

    return (
      <>
        {props.errors && <FormErrorSummary errors={props.errors} onJump={onJump} />}
        <Component {...props} bindFieldRef={bindRef} idPrefix={idPrefix} />
      </>
    );
  };
}
