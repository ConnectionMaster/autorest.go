/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutoRestExtension, } from '@azure-tools/autorest-extension-base';
import { namer } from './namer/namer';
import { transform } from './transform/transform';
import { generator } from './generator/generator';

require('source-map-support').install();

export async function main() {
  const pluginHost = new AutoRestExtension();
  pluginHost.Add('go', generator);
  pluginHost.Add('go-namer', namer);
  pluginHost.Add('go-transform', transform);
  await pluginHost.Run();
}

main();
