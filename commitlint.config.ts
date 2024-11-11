export default {
  extends: [
    '@commitlint/config-conventional',
  ],
  rules: {
    "type-enum": [2, 'always', [
        'build',
        'changelog',
        'chore',
        'ci',
        'docs',
        'feat',
        'fix',
        'perf',
        'refactor',
        'reapply',
        'revert',
        'style',
        'test',
      ],
    ],
  },
};
