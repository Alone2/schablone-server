from typing import Any, Dict, Optional

import httpx

from ...client import AuthenticatedClient
from ...models.template import Template
from ...types import Response


def _get_kwargs(
    template_id: int,
    *,
    client: AuthenticatedClient,
) -> Dict[str, Any]:
    url = "{}/template/get/{templateId}".format(client.base_url, templateId=template_id)

    headers: Dict[str, str] = client.get_headers()
    cookies: Dict[str, Any] = client.get_cookies()

    return {
        "method": "get",
        "url": url,
        "headers": headers,
        "cookies": cookies,
        "timeout": client.get_timeout(),
    }


def _parse_response(*, response: httpx.Response) -> Optional[Template]:
    if response.status_code == 200:
        response_200 = Template.from_dict(response.json())

        return response_200
    return None


def _build_response(*, response: httpx.Response) -> Response[Template]:
    return Response(
        status_code=response.status_code,
        content=response.content,
        headers=response.headers,
        parsed=_parse_response(response=response),
    )


def sync_detailed(
    template_id: int,
    *,
    client: AuthenticatedClient,
) -> Response[Template]:
    """Get a template

    Args:
        template_id (int):

    Returns:
        Response[Template]
    """

    kwargs = _get_kwargs(
        template_id=template_id,
        client=client,
    )

    response = httpx.request(
        verify=client.verify_ssl,
        **kwargs,
    )

    return _build_response(response=response)


def sync(
    template_id: int,
    *,
    client: AuthenticatedClient,
) -> Optional[Template]:
    """Get a template

    Args:
        template_id (int):

    Returns:
        Response[Template]
    """

    return sync_detailed(
        template_id=template_id,
        client=client,
    ).parsed


async def asyncio_detailed(
    template_id: int,
    *,
    client: AuthenticatedClient,
) -> Response[Template]:
    """Get a template

    Args:
        template_id (int):

    Returns:
        Response[Template]
    """

    kwargs = _get_kwargs(
        template_id=template_id,
        client=client,
    )

    async with httpx.AsyncClient(verify=client.verify_ssl) as _client:
        response = await _client.request(**kwargs)

    return _build_response(response=response)


async def asyncio(
    template_id: int,
    *,
    client: AuthenticatedClient,
) -> Optional[Template]:
    """Get a template

    Args:
        template_id (int):

    Returns:
        Response[Template]
    """

    return (
        await asyncio_detailed(
            template_id=template_id,
            client=client,
        )
    ).parsed
