/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { clientAdapter } from './clients.js';
import { typeAdapter } from './types.js';
import { GoEmitterOptions } from '../lib.js';
import * as go from '../../../codemodel.go/src/gocodemodel.js';
import { packageNameFromOutputFolder } from '../../../naming.go/src/naming.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { EmitContext } from '@typespec/compiler';

const headerText = `Copyright (c) Microsoft Corporation. All rights reserved.
Licensed under the MIT License. See License.txt in the project root for license information.
Code generated by Microsoft (R) Go Code Generator.`;

export function tcgcToGoCodeModel(context: EmitContext<GoEmitterOptions>): go.CodeModel {
  const info = new go.Info('TODO Title');
  const options = new go.Options(headerText, context.options['generate-fakes'] === true, context.options['inject-spans'] === true, context.options['disallow-unknown-fields'] === true);
  if (context.options['azcore-version']) {
    options.azcoreVersion = context.options['azcore-version'];
  }

  const codeModel = new go.CodeModel(info, 'data-plane', packageNameFromOutputFolder(context.emitterOutputDir), options);
  if (context.options.module) {
    codeModel.options.module = context.options.module;
  }
  if (context.options['module-version']) {
    codeModel.options.moduleVersion = context.options['module-version'];
  }
  if (context.options['rawjson-as-bytes']) {
    codeModel.options.rawJSONAsBytes = true;
  }
  if (context.options['slice-elements-byval']) {
    codeModel.options.sliceElementsByval = true;
  }

  const sdkContext = tcgc.createSdkContext(context);
  const ta = new typeAdapter(codeModel);
  ta.adaptTypes(sdkContext);

  const ca = new clientAdapter(ta, context.options);
  ca.adaptClients(sdkContext.experimental_sdkPackage);
  codeModel.sortContent();
  return codeModel;
}
