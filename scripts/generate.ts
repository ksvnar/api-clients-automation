import { callGenerator, setupAndGen } from './common.js';
import { getLanguageFolder } from './config.js';
import { formatter } from './formatter.js';
import { removeExistingCodegen } from './pre-gen/index.js';
import type { Generator } from './types.js';

async function preGen(gen: Generator): Promise<void> {
  await removeExistingCodegen(gen);
}

export async function generate(generators: Generator[]): Promise<void> {
  await setupAndGen(generators, 'client', async (gen) => {
    await preGen(gen);
    await callGenerator(gen);
  });

  for (const lang of [...new Set(generators.map((gen) => gen.language))]) {
    let folder = getLanguageFolder(lang);

    // We have scoped output folder for JavaScript which allow us to
    // avoid linting the whole client, only the part that changed
    if (lang === 'javascript') {
      folder = generators.reduce((folders, gen) => {
        if (gen.language === 'javascript') {
          return `${folders} ${gen.output}`;
        }

        return folders;
      }, '');
    }

    await formatter(lang, folder);
  }
}
