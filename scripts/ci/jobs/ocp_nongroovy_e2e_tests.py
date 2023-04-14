#!/usr/bin/env -S python3 -u

"""
Run tests/e2e in a openshift 4 cluster provided via a hive cluster_claim.
"""
import os
#from base_qa_e2e_test import make_qa_e2e_test_runner
from clusters import OpenShiftScaleWorkersCluster
from pre_tests import PreSystemTests
from ci_tests import NonGroovyE2e
from post_tests import PostClusterTest, FinalPost

# set required test parameters
os.environ["DEPLOY_STACKROX_VIA_OPERATOR"] = "true"
os.environ["ORCHESTRATOR_FLAVOR"] = "openshift"
os.environ["ROX_POSTGRES_DATASTORE"] = "true"
os.environ["ROX_ACTIVE_VULN_MGMT"] = "true"


ClusterTestRunner(
    cluster = OpenShiftScaleWorkersCluster("nongroovy-test"),
    pre_test=PreSystemTests(),
    test=NonGroovyE2e(),
    post_test=PostClusterTest(collect_central_artifacts=False),
    final_post=FinalPost(handle_e2e_progress_failures=False),
).run()